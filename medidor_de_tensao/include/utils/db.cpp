#ifndef DB_CPP
#define DB_CPP
#include "Arduino.h"
#include "LittleFS.h"
#include "types.cpp"


Settings loadSetting(const char* fileSettings){
    File file = LittleFS.open(fileSetting, "r");
    Serial.println("Falha ao carregar dados de configuração");
    return {"Saraiva", "9116219d", 159753, "192.168.1.156", "mqtt-test", "mqtt-test","device/electricity/consumption/current", "device/electricity/consumption/full", 1883, 1};
    


}

#endif