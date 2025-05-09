from datetime import timedelta, timezone, datetime
import subprocess
import orjson
import time 
import os

HEADER = '-' * 50
SHORT_HEADER = '-' * 25
REQUEST_FILENAME = 'request.json'
RESPONSE_FILENAME = 'response.json'

TEMP_ICON = '\U0001F321'
HUMIDITY_ICON = '\U0001F4A7'
CLOUD_ICON = '\U00002601'
CLOCK_ICON = '\U0001F55B'
WORLD_ICON = '\U0001F30D'
ERROR_ICON = '\U0000274C'


def main():
    locations = list()

    if not os.path.exists('.env'):
        print(HEADER)
        print('No .env file found\nPress enter to exit')
        input()
        return

    while True:
        print(HEADER)
        user_input = input('Enter a location: ')
        user_input = user_input.lower().strip().title()

        if not user_input:
            print('Please enter a location')
            continue
        elif user_input in locations:
            print('Location already exists')
            continue
        elif user_input.lower() in ('exit', 'break', 'stop', 'quit', 'q'):
            break

        locations.append(user_input)
        print(f'{user_input!r} added')
    
    if not locations:
        print('No locations given, exiting...')
        return
    
    locations = [{'name_city': location} for location in locations]

    with open(REQUEST_FILENAME, 'wb') as f:
        f.write(orjson.dumps({'Location': locations}))
        f.close()

    subprocess.Popen([
        'bin/backend-weather.exe',
        REQUEST_FILENAME]
    )

    while True:
        if os.path.exists(RESPONSE_FILENAME):
            break
        time.sleep(1)


    with open(RESPONSE_FILENAME, 'rb') as f:
        response = orjson.loads(f.read())
        f.close()

    print("\033c", end='')
    print(HEADER)
    for weather in response:
        print(f'Name: {weather["name"]}')

        if 'error' in weather:
            print(f'Error: {ERROR_ICON} {weather["error"]}')
        else:
            offset = weather["timezone"]

            timezone_offset = timezone(timedelta(seconds=offset))
            current_time = datetime.now(timezone_offset)
            current_time = current_time.strftime('%H:%M:%S %d/%m/%Y')

            print(
f'''Temprature:   {TEMP_ICON} {weather["main"]["temp"]}°C
Humidity:     {HUMIDITY_ICON}{weather["main"]["humidity"]}%
Description:  {CLOUD_ICON}  {weather["weather"][0]["description"].title()}

Timezone:     {WORLD_ICON} {weather["gmt"]}
Current time: {CLOCK_ICON} {current_time}
            ''')
        print(SHORT_HEADER)
    os.remove(REQUEST_FILENAME)
    os.remove(RESPONSE_FILENAME)
    print(f'\n{HEADER}\nPress enter to exit\n{HEADER}')
    input()


if __name__ == "__main__":
    main()
