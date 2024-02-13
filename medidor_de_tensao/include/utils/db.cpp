#ifndef DB_CPP
#define DB_CPP
#include "Arduino.h"
#include "LittleFS.h"



Settings loadSetting(const char* fileSettings){
    File file = LittleFS.open(fileSetting, "r");
    if(!file){
        Serial.println("Falha ao carregar dados de configuração");
        return {"Saraiva", "9116219d", 159753, "192.57.1.6", "device/electricity/consumption/full", "device/electricity/consumption/current", 1883, 5};
    }

    while(file.available()){
        
    }
}

#endif