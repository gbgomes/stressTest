# Execução e testes
## Sem uso do Docker
A partir da pasta raiz do projeto rodar o comando:

    go run main.go --url=<<URL a ser testada>> -r=<<Quantidade de requisições>> -c=<<Quantidade de requisições simultâneas>>

Exemplo

    go run main.go --url=http://www.google.com -r=10 -c=2

O comando acima irá executar 10 requisiões à URL www.google.com com 2 chamadas simultâneas

## Com uso do Docker
Fazer o build da imagem:

    docker build -t stresstest .

Executar com o comando:

    docker run stresstest --url=<<URL a ser testada>> -r=<<Quantidade de requisições>> -c=<<Quantidade de requisições simultâneas>> 

Exemplo

    docker run stresstest --url=http://www.google.com -r=10 -c=2

O comando acima irá executar 10 requisiões à URL www.google.com com 2 chamadas simultâneas

# Funcionamento
O programa Stresstest irá realizar um número determinado de chamadas, simultâneas ou não, a um servidor web expecífico, de acordo com os parâmetros informados.

Os parâmetros são os seguintes conforme o próprio help do programa:

    Usage:
    stresstest [flags]

    Flags:
    -c, --concurrency int   Número de chamadas simultâneas para a URL. Assume 10 se não for informado. (default 10)
    -h, --help              help for stressTest
    -r, --requests int      Número de requests a serem realizadas. Assume 50 se não for informado. (default 50)
    -u, --url string        Url a ser testada

Todos os parâmetros acima deverão ser informados via linha de comando, sendo que o parâmetro **url** é o único obrigatório. Os demais parâmetros, caso não informados assumirão os seguintes valores padrão:

    -r: 50
    -c: 10

Ao fim do processamento o seguinte relatório será exibido:

    URL: http://www.google.com.br, Requests: 50, Concurrency: 10

    Tempo total de execução: 278.460096ms

    Quantidade total de requests realizados: 50
    Quantidade total de requests com sucesso (http status 200): 10
    Quantidade total de requests com Erro: 40
    Distribuição dos erros:
     - Outros erros, Quantidade: 40

Sendo que os erros serão apresentados de acordo com o código do status de erro retornado pelo servidor.

## Exemplos de execução e seus resultados

**go run main.go --url=http://www.google.com.br -r=10 -c=2**

    URL: http://www.google.com.br, Requests: 10, Concurrency: 2

    Tempo total de execução: 164.935524ms

    Quantidade total de requests realizados: 10
    Quantidade total de requests com sucesso (http status 200): 2
    Quantidade total de requests com Erro: 8
    Distribuição dos erros:
      - Outros erros, Quantidade: 8


**go run main.go --url=http://www.vivo.com.br -r=10 -c=2**

    URL: http://www.vivo.com.br, Requests: 10, Concurrency: 2

    Tempo total de execução: 207.897366ms

    Quantidade total de requests realizados: 10
    Quantidade total de requests com sucesso (http status 200): 0
    Quantidade total de requests com Erro: 10
    Distribuição dos erros:
      - Erro 403, Quantidade: 2
      - Outros erros, Quantidade: 8