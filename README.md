Este projeto tem como objetivo fazer envio de email 

para configuralo é bem simples basta criar um arquivo .env na raiz de seu projeto e neste arquivo ira colocar as seguintes variaveis

KEY = "*******" - Aqui ira colocar a chave/senha do serviço que ira utilizar no caso para fins de teste você pode utilizar o proprio Gmail que irei mostrar mais a baixo como faz para pegar

FROM = "exemplo@gmail.com" - Aqui ira colocar quem estará enviando o email

TO = "exemplo@gmail.com" - Aqui ira infomrar quem estará recebendo esse email

PORT = 465 - Aqui ira informar a porta

HOST = "smtp.gmail.com" servidor smtp que ira fazer o envio caso va utilizar do gmail pode usar como descrito aqui


Para Utilizar sua propria conta google para configurar e fazer o envio de emails atraves dela basta estar seguinto estes passos:

1 - Acesse sua conta google e va em segurança
2- Verifique se esta com a autenticação de 2 fatores ativado caso não esteja ative 
3 - na barra de pesquisa digite "Senhas de app"
4 - De um nome para essa senha 
5 - Apos isso ira aparecer um codigo em sua tela esse codigo sera sua KEY

espero ter conseguido explicar como configurar caso tenha qualquer duvida pode estar entrando em contato comigo
