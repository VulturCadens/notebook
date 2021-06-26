#define BEGIN_CHAR 0x5B
#define END_CHAR   0x5D

const uint16_t MASK = B11111111;

uint8_t sendArray[4];

uint16_t analogValue = 0;
uint16_t currentValue = 0;

char c;
boolean isCodeReady;
String code;

void setup() {
    Serial.begin(9600);

    sendArray[0] = BEGIN_CHAR;
    sendArray[3] = END_CHAR;

    pinMode(LED_BUILTIN, OUTPUT);
    digitalWrite(LED_BUILTIN, LOW);
}

void loop() {
    if (Serial.available() > 0) {
        c = Serial.read();

        switch (c) {
            case BEGIN_CHAR:
                code = "";
                break;

            case END_CHAR:
                isCodeReady = true;
                break;

            default:
                code.concat(c);
                break;
        }
    }

    if (isCodeReady) {
        if (code == "ON") {
            digitalWrite(LED_BUILTIN, HIGH);
        }

        if (code == "OFF") {
            digitalWrite(LED_BUILTIN, LOW);
        }

        isCodeReady = false;
    }

    analogValue = analogRead(A0);

    if (analogValue != currentValue) {
        currentValue = analogValue;

        sendArray[1] = analogValue >> 8;
        sendArray[2] = analogValue & MASK;

        Serial.write(sendArray, 4);
    }

    delay(20);
}
