## Conceitos

## Descrição

## Start

## EndPoints

Metodo: GET
Rota: api/v1/candieiro/points
Querys: 
    plantId: number
Retornos: [{id int, name string, lastConsumer float}]

Metodo: GET
Rota: api/v1/candieiro/point/consumption
Parâmetros:
    startMoment: Date
    endMoment: Date
    pointId: Int
Retornos: [{moment Date, value float kW float}]

Metodo: GET
Rota: api/v1/candieiro/point/shutdowns
Parâmetros:
    pointId: Int
Retornos: [{startShutdown Date, endShutdown Date}]

Metodo: GET
Rota: api/v1/candieiro/point/alert
Parâmetros:
    pointId: Int
Retornos: {alertLimit float}

Metodo: POST
Rota: api/v1/candieiro/point/shutdown
Body:
    pointId: Int
    startShutdown: Date
    endShutdonw: Date
Retornos: {startShutdown Date, endShutdown Date}

Metodo: DELETE
Rota: api/v1/candieiro/point/shutdown
Query:
    shutdownId: Int
Retornos: 

Metodo: PUT
Rota: api/v1/candieiro/point/alert
Parâmetros:
    pointId: Int
    alertLimit: float
Retornos: {alertLimit float}



## Contrato de interface

### Dispositivo

O dispositivo de medição de consumo de energia deve realizar o envio para o broker das atualizações a cada vez que a potência consumida variar. As atualizações devem ter a seguinte estrutura:

{
    kw/h: número representando o consumo atual
    start: o timestamp que representa o inicio desse período de consumo
    end: o timestamp que representa o fim desse período de consumo ou 0 caso o período não tenha finalizado
}