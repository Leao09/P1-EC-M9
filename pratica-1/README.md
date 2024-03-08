# Ponderada - Testes no simulador MQTT

## Como rodar o projeto 
- É necessário primeiramente clonar o repositório utilizando o seguinte comando:
<pre><code>
 git clone  https://github.com/Leao09/Exercicios-prog-M9.git
</code></pre>
- Depois entrar no diretório do projeto
<pre><code>
 cd ponderada-2
</code></pre> 
- Em sequida é necessário ter instalado a linguagem  go e o broker mosquitto para executar o projeto
- Esses arquivos podem ser baixados pelos seguintes links [Go](https://go.dev/dl/) e [Mosquitto](https://mosquitto.org/download/)
- Assim, inicie um módulo para o diretório e depois baixe as depencias para o projeto
<pre><code>
 go mod init seu-no
 go mod tidy
</code></pre>
- Agora é necessário rodar um comando para ativar o brocker localmente 
<pre><code>
 mosquitto -c mosquitto.conf 
</code></pre> 
- Com o brocker funcionando, entramos na página publisher e rodamos um comando para ativá-lo
<pre><code>
cd publisher
 go run publisher.go
</code></pre> 
- Por fim, podemos rodar os tests que irão fazer validações sobre o nosso publisher 
<pre><code>
go test -v -cover 
</code></pre> 
## Video 
[Link](https://youtu.be/s7dnlcrvFvc)