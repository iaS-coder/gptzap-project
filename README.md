# gptzap-project
A ideia do projeto 'e realizar a integracao entre o chatgpt e o whatsapp.

A api em questao 'e uma api serverless, que roda em um lambda hospedado na aws, esse lamda recebe uma requisicao via um endpoint.

Apartir do momento em que a mensagem 'e enviada do whatsapp utilizando o Twillio como intermediador, ele possui uma conexao com o endpoint do Gateway desse lambda, realizando a request o lambda chama a API do chatgpt com a pergunta informada na mensagem do Whatsapp, apos o recebimento do retorno, essa mensagem volta para o chat do Whatsapp
