# Exercícios Diagnósticos

## Exercícios Simples

### Classe "Círculo"

1. **Objetivo:** Desenvolver uma classe `Círculo` com um atributo de raio.  
   **Tarefas:**  
   - Implemente métodos para calcular a área e o perímetro do círculo.
   - Crie um método adicional para exibir as informações do círculo, incluindo raio, área e perímetro.

### Classe "Conta Bancária"

2. **Objetivo:** Implementar uma classe `Conta Bancária` com os seguintes atributos: número da conta, nome do titular e saldo.  
   **Tarefas:**  
   - Desenvolva métodos para realizar depósitos, saques e para exibir o saldo atual da conta.

### Classe "Pessoa"

3. **Objetivo:** Criar uma classe `Pessoa` que contenha os atributos nome, idade e gênero.  
   **Tarefas:**  
   - Implemente métodos para exibir as informações da pessoa.
   - Adicione um método para verificar se a pessoa é maior de idade (idade maior ou igual a 18 anos).

### Classe "Retângulo"

4. **Objetivo:** Desenvolver uma classe `Retângulo` com atributos para largura e altura.  
   **Tarefas:**  
   - Implemente métodos para calcular a área e o perímetro do retângulo.
   - Adicione um método para exibir as dimensões do retângulo.

## Exercícios Médios

### Classe "Livro"

5. **Objetivo:** Criar uma classe `Livro` com os seguintes atributos: título, autor, ano de publicação e número de páginas.  
   **Tarefas:**  
   - Desenvolva métodos para emprestar e devolver o livro.
   - Implemente um método para verificar a disponibilidade do livro para empréstimo.

### Classe "Lâmpada"

6. **Objetivo** Criar uma classe `Lâmpada` a qual pode ser ligada e desligada. Também deve ser possível observar o estado da lâmpada (se desligada ou ligada).
   **Tarefas**
   - Desenvolva uma nova classe para a lâmpada de forma a incluir as características de potência e voltagem. Garanta que seja possível tanto ler quanto alterar os valores de potência e voltagem de uma lâmpada.
   - Crie uma classe Teste com um método main para testar as classes desenvolvidas nos exercícios 1 e 2. Crie uma lâmpada, apresente no console as informações de estado (se ligada ou desligada, potência e voltagem), ligue a lâmpada e apresente novamente as informações de estado.
   - Modifique a classe da lâmpada criada anteriormente para incluir o caso de uma lâmpada queimar ao ser ligada. Sabe-se que existe uma chance de 15% da lâmpada queimar ao ser ligada. Dica: neste exercício é importante pesquisar na biblioteca de classes fornecida pela linguagem de programação uma classe que dê suporte à geração de números aleatórios


### Exercícios Difíceis

### Classe "Conta Bancária"

7. **Objetivos** Faça uma classe `Conta` que contenha o nome do cliente, o número da conta e o saldo. 
   **Tarefas**
   - Estes valores deverão ser informados no construtor da classe. Faça um método depositar e um método sacar para realizar o depósito e saque de valores da conta. 
   - Faça um método obterSaldo que retorne o saldo do cliente. 
   - Faça um método obterNumero que retorne o número da conta. 
   - Faça um método obterNomeCliente que retorna o nome do cliente titular da conta.
   - Desenvolva um programa para testar a classe Conta. O programa deverá criar 3 contas diferentes e realizar diversas operações de saque e depósito. Depois, o programa deverá emitir um relatório (no console) apresentando o número, o titular e o saldo atual de cada conta.
   - Faça uma classe Extrato, que contenha a data da movimentação e o valor movimentado (pode ser tanto positivo quanto negativo);
   - Incremente a classe Conta, incluindo um vetor de 1000 posições de Extrato.
   - Cada vez que ocorrer um depósito ou um saque na conta, deverá criar um objeto Extrato com as informações de data da movimentação e o valor da movimentação;
   - Desenvolva um programa para testar a nova classe Conta. O programa deverá criar 3 contas diferentes e realizar diversas operações de saque e depósito. Depois, o programa deverá emitir um relatório (no console) apresentando o número, o titular, saldo atual e o extrato de cada conta.
   - Faça uma classe Cartão, que contenha o número do cartão, a validade do cartão.
   - Incremente a classe Conta, incluindo a informação de cartão.
   - Desenvolva um programa para testar a nova classe Conta. O programa deverá criar 3 contas diferentes e realizar diversas operações de saque e depósito. O programa deverá realizar operações de saque com cartão, onde deverá solicitar o número e a validade do cartão e identificar a conta à qual pertence o cartão. Depois, o programa deverá emitir um relatório (no console) apresentando o número, o titular, saldo atual e o extrato de cada conta.

