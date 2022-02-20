#define BEGIN_CHAR 0x5B // [
#define END_CHAR 0x5D   // ]
#define MAX_STRING 10

uint8_t sendArray[3];

uint16_t analogValue = 0;
uint16_t currentValue = 0;

char c;
char code[MAX_STRING];
char s[2];

bool isCodeReady = false;

void setup()
{
    Serial.begin(19200);

    pinMode(LED_BUILTIN, OUTPUT);
    digitalWrite(LED_BUILTIN, HIGH);

    s[1] = '\0'; // NULL character.
}

void loop()
{
    
    while (Serial.available() > 0) {
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

        // Return value 0 -> the contents of both strings are equal.
        // Numbers that are not equal to 0 are viewed as true in C.

        if (!strcmp(code, "HIGH")) {
            digitalWrite(LED_BUILTIN, HIGH);
        }

        if (!strcmp(code, "LOW")) {
            digitalWrite(LED_BUILTIN, LOW);
        }

        isCodeReady = false;
    }

    analogValue = analogRead(A0);

    if (analogValue != currentValue) {
        currentValue = analogValue;

        sendArray[0] = analogValue >> 8; // MSB
        sendArray[1] = analogValue & 0xFF; // LSB
        sendArray[2] = '\n'; // Newline character

        Serial.write(sendArray, 3);
    }

    delay(100);
}