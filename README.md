# Arara

**Arara** é uma ferramenta desenvolvida em Go para plataformas populares como **Salesforce**, **ServiceNow**, e outras APIs REST. A ferramenta foi projetada para ser leve, rápida e intuitiva, permitindo que você realize testes de forma eficiente.

O objetivo do projeto é oferecer uma abordagem leve, determinística e logicamente estruturada para identificação de ambientes corporativos que utilizam integrações específicas.

Arara opera como um sistema de verificação de hipóteses. A entrada fornecida pelo usuário é utilizada para instanciar padrões determinísticos que representam possíveis integrações SaaS. Cada padrão gera uma hipótese sobre a existência de um endpoint associado à organização analisada. A requisição HTTP funciona como operador de teste dessa hipótese. A resposta retornada pelo servidor define o estado lógico do recurso consultado. Um código 200 indica existência e acessibilidade. Um código 401 indica existência com controle de autenticação. Códigos 404 indicam inexistência do recurso naquele domínio. Erros de conexão indicam ausência ou indisponibilidade da infraestrutura correspondente.

A arquitetura interna utiliza concorrência estruturada por meio de goroutines, canais e sincronização com WaitGroup. O sistema implementa um pipeline composto por produtor, conjunto de trabalhadores e consumidor final. O produtor gera URLs a partir dos padrões conhecidos. Os trabalhadores executam requisições HTTP utilizando um cliente com tempo limite configurado. O consumidor recebe os resultados e realiza a classificação semântica com base no código de status. Esse modelo segue o paradigma de comunicação por canais inspirado em Communicating Sequential Processes, garantindo isolamento de responsabilidades e evitando condições de corrida.

A complexidade assintótica do algoritmo é linear em relação ao número de padrões definidos. Se n representa a quantidade de endpoints gerados, o sistema executa exatamente n requisições. Como a execução ocorre de forma concorrente com um número fixo de trabalhadores, o tempo total tende a aproximar-se de n dividido pelo número de workers, acrescido da latência média de rede. Esse comportamento torna o desempenho previsível e controlável.

O programa segue uma ideia matemática muito direta:
f(x) → (y, e)

Para utilizar a ferramenta, basta clonar o repositório, acessar o diretório do projeto e executar o arquivo principal com go run arara.go. Após informar o nome da empresa alvo, o sistema iniciará automaticamente a geração das URLs e a verificação concorrente dos endpoints associados.

Arara demonstra domínio de modelagem funcional, uso correto de concorrência estruturada em Go, tratamento explícito de erros, compreensão do protocolo HTTP e aplicação de raciocínio matemático simples
