#ifndef WIFI_CPP
#define WIFI_CPP

#include "Arduino.h"
#include "PubSubClient.h"
#include <WiFi.h>

WiFiClient espClient;
PubSubClient clientMqtt(espClient);


boolean startWiFi(const char* ssid, const char *password, unsigned long *lastConnection, unsigned long connectionInterval){
    if(WiFi.status() == WL_CONNECTED || (millis() - *(lastConnection)) < connectionInterval ) return WiFi.status() == WL_CONNECTED;
    WiFi.begin();

    return WiFi.status() == WL_CONNECTED;
}

void setCommunicationBroker(const char * ip, int port, MQTT_CALLBACK_SIGNATURE){
    clientMqtt.setServer(ip, port);
}

boolean reconnectBroker(const char *ssid, const char *password, const char* idMqtt, unsigned long *lastConnection){
    if(startWiFi(ssid, password, lastConnection, 30000)) return false;

    if(!clientMqtt.connected()){
        if(clientMqtt.connect(idMqtt)) Serial.println("Conectado.");
        else return false;
    }

    clientMqtt.loop();
    return true;
}

boolean sendMessageToServer(String message, String topic){
    return clientMqtt.publish(topic.c_str(), message.c_str());
}

#endif