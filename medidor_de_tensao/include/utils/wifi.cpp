#ifndef WIFI_CPP
#define WIFI_CPP

#include "Arduino.h"
#include "PubSubClient.h"
#include <ESP8266WiFi.h>

WiFiClient espClient;
PubSubClient clientMqtt(espClient);


boolean startWiFi(const char* ssid, const char *password, unsigned long *lastConnection, unsigned long connectionInterval){
    if(WiFi.status() == WL_CONNECTED || (millis() - *(lastConnection)) < connectionInterval ) return WiFi.status() == WL_CONNECTED;
    WiFi.begin(ssid, password);
    *(lastConnection) = millis();

    return WiFi.status() == WL_CONNECTED;
}

void setCommunicationBroker(const char * ip, int port, MQTT_CALLBACK_SIGNATURE){
    clientMqtt.setServer(ip, 1883);
    clientMqtt.setCallback(callback);
}

boolean reconnectBroker(const char *ssid, const char *password, const char *userBroker, const char *passwordBroker, int idMqtt, unsigned long *lastConnection){
    if(!startWiFi(ssid, password, lastConnection, 30000)) return false;
    
    if(!clientMqtt.connected()){
        if(!clientMqtt.connect(String(idMqtt).c_str(), userBroker, passwordBroker)) return false;
        Serial.println("Conectado.");
    }

    clientMqtt.loop();
    return true;
}

boolean sendMessageToServer(String message, String topic){
    Serial.println(message);
    Serial.println(topic);
    return clientMqtt.publish(topic.c_str(), message.c_str());
}

#endif