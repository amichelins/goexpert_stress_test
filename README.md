# goexpert_stress_test
Objetivo: Testar a carga de uma url utilizando um numero predefinido de requisições,
          pudendo indicar a quantidade de conconrrencia utilizada.
          Para este fim utilizamos o pkg cobra


Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests. Se não informado será utilizado o valor 100
--concurrency: Número de chamadas simultâneas.  Se não informado será utilizado o valor 2

Após realizadas as requisições estipuladas será mostrado um relatário com os dados coletados.
Relatório neste formato:

<pre>
------------------------------------------------------------------
                    STRESS TEST REPORT
------------------------------------------------------------------
-------------
Dados Gerais:
-------------
    Url: {URL}
    Concorrencia utilizada: {concorrencia utilizada}
    Tempo Gasto: {tempo gasto em segundos}
    Requisições Feitas: {numero de requisições realizadas}
    Requisições completadas (http 200): {Numero de requisições com http code 200}

------------------------
Dados por Códigos http:
------------------------
   1xx Informativos       ---> {Quantidade Http code 1xx}
   3xx Redireccionamentos ---> {Quantidade Http code 2xx}
   4xx Erro no Cliente    ---> {Quantidade Http code 4xx}
   5xx Erro no Servidor   ---> {Quantidade Http code 5xx}

Requisições com Erro : {Quantidade de requisições que deram erro}
</pre>


Para rodar o sistema temos duas formas de rodar:


    1 - Fazendo o build do projeto:
        go build -ldflags="-s -w" -o stress_tester .
        ./stress_tester

        ou

        go install .
        ./stress_tester

    2 - Via docker
        docker build --no-cache -t amichelin/stress:1.0 .
        docker run -it amichelin/stress:1.0 --url={url} --requests={Quantidade de request} --concurrency={Quantidade de concorrencia}



EXEMPLOS utilizando docker run:

1 - docker run -it amichelin/stress:1.0 --url=http://www.webcontinental.com.br --requests=1000

Resposta:
<pre>
------------------------------------------------------------------
                    STRESS TEST REPORT
------------------------------------------------------------------
-------------
Dados Gerais:
-------------
    Url: http://www.webcontinental.com.br
    Concorrencia utilizada: 2
    Tempo Gasto: 23.015892266s
    Requisições Feitas: 1000
    Requisições completadas (http 200): 1000

------------------------
Dados por Códigos http:
------------------------
   1xx Informativos       ---> 0
   3xx Redireccionamentos ---> 0
   4xx Erro no Cliente    ---> 0
   5xx Erro no Servidor   ---> 0

Requisições com Erro : 0
</pre>

    2 - docker run -it amichelin/stress:1.0 --url=http://www.webcontinental.com.br --requests=1000 --concurrency=10

    Resposta:
<pre>
------------------------------------------------------------------
                    STRESS TEST REPORT
------------------------------------------------------------------
-------------
Dados Gerais:
-------------
    Url: http://www.webcontinental.com.br
    Concorrencia utilizada: 10
    Tempo Gasto: 12.229890008s
    Requisições Feitas: 1000
    Requisições completadas (http 200): 1000

------------------------
Dados por Códigos http:
------------------------
   1xx Informativos       ---> 0
   3xx Redireccionamentos ---> 0
   4xx Erro no Cliente    ---> 0
   5xx Erro no Servidor   ---> 0

Requisições com Erro : 0
</pre>
Na pasta do projeto tem duas imagens evidenciando os resultados