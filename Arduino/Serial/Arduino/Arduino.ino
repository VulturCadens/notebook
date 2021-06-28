#define BEGIN_CHAR 0x5B
#define END_CHAR 0x5D
#define MAX_STRING 10

uint8_t sendArray[4];

uint16_t analogValue = 0;
uint16_t currentValue = 0;

char c;
char code[MAX_STRING];
char s[2];

bool isCodeReady = false;

void setup()
{
    Serial.begin(9600);

    sendArray[0] = BEGIN_CHAR;
    sendArray[3] = END_CHAR;

    pinMode(LED_BUILTIN, OUTPUT);
    digitalWrite(LED_BUILTIN, LOW);

    s[1] = '\0'; // NULL character.
}

void loop()
{
    if (Serial.available() > 0) {
        c = Serial.read();

        switch (c) {
            case BEGIN_CHAR:
                code[0] = '\0';
                break;

            case END_CHAR:
                isCodeReady = true;
                break;

            default:
                s[0] = c;

                if (strlen(code) != MAX_STRING - 1) {
                    strcat(code, s);
                }

                break;
        }
    }

    if (isCodeReady) {
        if (strcmp(code, "ON")) {
            digitalWrite(LED_BUILTIN, HIGH);
        }

        if (strcmp(code, "OFF")) {
            digitalWrite(LED_BUILTIN, LOW);
        }

        isCodeReady = false;
    }

    analogValue = analogRead(A0);

    if (analogValue != currentValue) {
        currentValue = analogValue;

        sendArray[1] = analogValue >> 8; // MSB
        sendArray[2] = analogValue & 0xFF; // LSB

        Serial.write(sendArray, 4);
    }

    delay(20);
}
