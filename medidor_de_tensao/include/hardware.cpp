#ifndef START_HARDWARE
#define START_HARDWARE

#include "Arduino.h"
#include "LittleFS.h"
#include "ACS712.h"

ACS712 sensor(ACS712_20A, 4);

void startHardware(boolean calibrate){
    if(!LittleFS.begin(true)){
        Serial.println("Falha ao abrir sistema de arquivos");
        ESP.restart();
    }

    if(calibrate)sensor.calibrate();
}

float readAcs712(){
    return sensor.getCurrentAC();
}


#endif