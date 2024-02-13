#ifndef CORE_CPP
#define CORE_CPP

#include "Arduino.h"
#include "types.cpp"
#include "utils/mqtt.cpp"

boolean updatePowerOnServer(UpdateErrorCallback callback, float power){
}

boolean sendMeasurementRecord(UpdateErrorCallback callback, ConsumptionRecord consumption){

}

float calculateKwh(float current){
    return current * 3600;
}


ConsumptionRecord resetMeasurementPeriod(unsigned long currentTime){
    return {currentTime, 0, 0, 0};
}


boolean validateMeasurementPeriod(ConsumptionRecord consumption, unsigned long time, int tempoMaximo){
    return (time - consumption.startConsumption)/60 < tempoMaximo;
}

ConsumptionRecord recordsMeasurementInPeriod(ConsumptionRecord consumption, float power){
    consumption.fullPower += power;
    return consumption;
}

#endif