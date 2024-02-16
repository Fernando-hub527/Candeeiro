#ifndef CORE_CPP
#define CORE_CPP

#include "Arduino.h"
#include "types.cpp"
#include "utils/mqtt.cpp"

String makeMessageCurrentConsumption(int kwh, unsigned long time, int serialNumber){
  return "{\"kw_h\": "+ String(kwh) +", \"time\": "+ time +", \"serialNumber\": "+ serialNumber +"}";
}

String makeMessageConsumption(ConsumptionRecord consumption, int serialNumber){
  return "{\"kw\": "+ String(consumption.fullPower) +", \"startTime\": "+ consumption.startConsumption +", \"endTime\": "+ consumption.endConsumption +", \"serialNumber\": "+ serialNumber +"}";
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