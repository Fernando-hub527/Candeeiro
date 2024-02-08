#ifndef WIFI_CPP
#define WIFI_CPP

#include "Arduino.h"
#include "PubSubClient.h"
#include <WiFi.h>

WiFiClient espClient;
PubSubClient clientMqtt(espClient);


boolean startWiFi(const char* ssid, const char *password){
    if(WiFi.status() == WL_CONNECTED) return true;

    WiFi.begin();
    while (WiFi.status() != WL_CONNECTED){
        delay(10);
    }

    return true;
    
}

boolean setCommunicationBroker(const char * ip, int port, MQTT_CALLBACK_SIGNATURE){
    clientMqtt.setServer(ip, port);
    clientMqtt.setCallback(callback);
}

#endif