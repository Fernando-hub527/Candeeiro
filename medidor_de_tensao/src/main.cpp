#include <Arduino.h>
#include "core.cpp"
#include "utils/db.cpp"
#include "hardware.cpp"
#include "utils/wifi.cpp"
#include "utils/callbacks.cpp"
#include "utils/time.cpp"

const int volts = 230;


void setup() {
  Serial.begin(9600);
  startHardware(true);
  settings = loadSetting(fileSetting);
  setCommunicationBroker(settings.ipBroker, settings.brokerDoor, callbackBroker);
  startWiFi(settings.ssid, settings.password, &lastConnection, 0);
  startTime();
}

String makeMessageCurrentConsumption(int kwh, unsigned long time, int serialNumber){
  return "{\"kw_h\": "+ String(kwh) +", \"time\": "+ time +", \"serialNumber\": "+ serialNumber +"}";
}

String makeMessageConsumption(ConsumptionRecord consumption, int serialNumber){
  return "{\"kw\": "+ String(consumption.fullPower) +", \"startTime\": "+ consumption.startConsumption +", \"endTime\": "+ consumption.endConsumption +", \"serialNumber\": "+ serialNumber +"}";
}

void loop() {
  float power = 220 * readAcs712();

  sendMessageToServer(makeMessageCurrentConsumption(calculateKwh(power), 1, settings.serialNumber), settings.currentConsumption);

  if (!validateMeasurementPeriod(consumptionRecord, getCurrentTime(), settings.measurementTime)){
    sendMessageToServer(makeMessageConsumption(consumptionRecord, settings.serialNumber), settings.consumptionPeriod);
    consumptionRecord = resetMeasurementPeriod(getCurrentTime());
  }

  recordsMeasurementInPeriod(consumptionRecord, power);

}
