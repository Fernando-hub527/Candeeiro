#ifndef TYPES_CPP
#define TYPES_CPP
#include "Arduino.h"

typedef void UpdateErrorCallback(String, String);

struct ConsumptionRecord{
    unsigned long startConsumption;
    unsigned long endConsumption;
    int fullPower;
    int kwh;
} consumptionRecord;

struct Settings{
    const char *ssid;
    const char *password;
    unsigned int serialNumber;
    const char *ipBroker;
    const char *currentConsumption;
    const char *consumptionPeriod;
    int brokerDoor;
    int measurementTime;
}settings;

const char* fileSetting = "fileSettings.txt";
unsigned long lastConnection = 0;

#endif