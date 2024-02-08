#include <Arduino.h>
#include "ACS712.h"
#include "core.cpp"
#include "utils/db.cpp"

ACS712 sensor(ACS712_20A, 4);
const int volts = 230;


void setup() {
  Serial.begin(9600);
  sensor.calibrate();

}

void loop() {
  float power = getPower();
  updatePowerOnServer(recordError, power);
  if (validateMeasurementPeriod()){
    sendMeasurementRecord(recordError);
    resetMeasurementPeriod();
  }

}
