#include <Arduino.h>
#include "core.cpp"
#include "utils/db.cpp"
#include "hardware.cpp"
#include "utils/wifi.cpp"

const int volts = 230;


void setup() {
  Serial.begin(9600);
  startHardware();
  settings = loadSetting(fileSetting);
  setCommunicationBroker(settings.ipBroker, settings.brokerDoor);
  startWiFi(settings.ssid, settings.password);

}

void loop() {
  float power = getPower();
  updatePowerOnServer(recordError, power);

  if (validateMeasurementPeriod(consumptionRecord)){
    sendMeasurementRecord(recordError, consumptionRecord);
    resetMeasurementPeriod(consumptionRecord);
  }
  recordsMeasurementInPeriod(consumptionRecord);

}
