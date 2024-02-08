#ifndef TYPES_CPP
#define TYPES_CPP
#include "Arduino.h"

typedef void(*UpdateErrorCallback)(String, String);

struct ConsumptionRecord{
    unsigned long startConsumption;
    unsigned long endConsumption;
    int fullPower;
    int kwh;
} consumptionRecord;

struct Settings{
    const char *ssid;
    const char *password;
    const char *serialNumber;
    const char *ipBroker;
    int brokerDoor;
    int measurementTime;
}settings;

const char* fileSetting = "fileSettings.txt";

#endif