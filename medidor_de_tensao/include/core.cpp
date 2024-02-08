#ifndef CORE_CPP
#define CORE_CPP

#include "Arduino.h"
#include "types.cpp"

int getPower(){

}

boolean updatePowerOnServer(UpdateErrorCallback callback, float power){

}

boolean sendMeasurementRecord(UpdateErrorCallback callback, ConsumptionRecord consumption){

}

void resetMeasurementPeriod(ConsumptionRecord consumption){

}


boolean validateMeasurementPeriod(ConsumptionRecord consumption){

}

void recordsMeasurementInPeriod(ConsumptionRecord consumption){
    
}

#endif