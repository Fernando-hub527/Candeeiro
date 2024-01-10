# candieiro
Sistema voltado para o controle do consumo de energia, permitindo o monitoramento em tempo real e análise de histórico

## Definição de requisitos

O sistema é formado por três componentes que trabalham em conjunto para permitir que o objetivo descrito anteriomente seja alcançado. Esses componentes são:

* Dispositivo
    </br>O disposito tem como responsibilidade coletar o consumo de determinado equipamento e realizar o envio dessa informação para o servidor. Ele deve também realizar desligamentos programados em determinados períodos.
* Api
    </br>A api tem como responsabilidade ser o intermediário entre os clientes e os pontos de consumo, precisando assim disponibilizar um interface para cada. 
    </br>A interface com o dispositivo deve ser capaz de realizar o processamento dos dados recebidos assim como o envio de comandos de desligamento.
    </br>A interface com cliente deve disponibilizar meios para que o mesmo acesse os registros de consumo dos pontos que lhe forem permitidos assim como a configuração dos mesmos

* Front end
    </br>O front end tem como responsabilidade disponibilizar meios para que o cliente final visualize o consumo das plantas que lhe pertencerem, permitindo a análise do histórico de consumo e o consumo em tempo real da sua planta. Os clientes devem ser capazes de criar plantas e pontos, assim como realizar a configuração dos mesmos

## Contratos de interfaces:

### Dispositivos:

* Os dispositivos devem realizr o envio do consumo atual e do consumo por período, sendo que o consumo atual deve ser enviado sempre que houver variação no consumo, enquanto o consumo por período deve ser enviado de acordo intervalo de tempo configurado. Além dessas informações o equipamento deve informar sobre seu estado de funcionamento. Os dados devem seguir o seguinte padrão:
</br>

1. Consumo atual:</br>
**{"kw_h": number, "time": number, "deviceId": number}**
</br>

2. Consumo por intervalo:</br>
**{"kw": number, "startTime": number, "endTime": number, "deviceId": number}**

3. Funcionamento: <br>
**{"on": boolean, "time": number, "deviceId": number}**

* Além das informações enviadas o dispositivo também aceita comandos de desligamento na seguinte estrutura:

1. Desligamentos: </br>
**{"startTime": number, "endTime": number, "recurrent": boolean}**

### Api:

* A interface da api com o dispositivo deve respeitar o que for definido pelo dispositivo

* A interface com o cliente deve ser feita através de métodos http, onde será possível acessar os recursos da aplicação, e via websocket, de forma que seja possível acessar os dados em tempo real, as interfaces terão a seguinte estrutura: <br>

1. Listagem de pontos 
    - Metodo: GET
    - Rota: api/v1/candieiro/plant/points
    - Querys: 
        - plantId: number
    - Retornos: [{id: int, name: string, lastConsumer: float}]

2. Listagem de consumo por ponto e intervalo
    - Metodo: GET
    - Rota: api/v1/candieiro/plant/point/consumption
    - Querys:
        - startMoment: Date
        - endMoment: Date
        - pointId: Int
    - Retornos: [{startMoment: Date, endMoment: Date, value: float, kW: float}]

3. Listagem de desligamentos por ponto
    - Metodo: GET
    - Rota: api/v1/candieiro/plant/point/shutdowns
    - Querys:
        - pointId: Int
    - Retornos: [{startShutdown: Date, endShutdown: Date}]

4. Cadastro de desligamento
    - Metodo: POST
    - Rota: api/v1/candieiro/plant/point/shutdown
    - Body:
        - pointId: Int
        - startShutdown: Date
        - endShutdonw: Date
    - Retornos: {startShutdown: Date, endShutdown: Date}

5. Cadastro de usuário
    - Metodo: POST
    - Rota: api/v1/candieiro/user
    - Body:
        - userName: string
        - password: string
        - email: string

6. Cadastro de planta
    - Metodo: POST
    - Rota: api/v1/candieiro/plant
    - Body:
        - name: string

7. Cadastro de ponto de consumo
    - Metodo: POST
    - Rota: api/v1/candieiro/plant/point
    - Body:
        - plantId: number

8. Remoção de desligamento
    - Metodo: DELETE
    - Rota: api/v1/candieiro/point/shutdown
    - Query:
        - shutdownId: Int

9. Login
