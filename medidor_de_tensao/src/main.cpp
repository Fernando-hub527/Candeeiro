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
  startWiFi(settings.ssid, settings.password, &lastConnection, 0);

  setCommunicationBroker(settings.ipBroker, settings.brokerDoor, callbackBroker);
  startTime();
  Serial.println("Iniciando");
  consumptionRecord = resetMeasurementPeriod(getCurrentTime());
}



void loop() {
  delay(1000);
  reconnectBroker(settings.ssid, settings.password, settings.useBroker, settings.passwordBroker,  settings.serialNumber, &lastConnection);
  float power = 220 * readAcs712();

  sendMessageToServer(makeMessageCurrentConsumption(calculateKwh(power), getCurrentTime(), settings.serialNumber), settings.currentConsumption);

  if (!validateMeasurementPeriod(consumptionRecord, getCurrentTime(), settings.measurementTime)){
    sendMessageToServer(makeMessageConsumption(consumptionRecord, settings.serialNumber), settings.consumptionPeriod);
    consumptionRecord = resetMeasurementPeriod(getCurrentTime());
  }

  recordsMeasurementInPeriod(consumptionRecord, power);

}
