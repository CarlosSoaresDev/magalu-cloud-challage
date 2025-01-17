How run in (magalu cloud) https://console.magalu.cloud/
```markdown
## How to Run in Magalu Cloud

To run your application in Magalu Cloud, follow these steps:

1. **Login to Magalu Cloud Console**: Open your browser and go to [Magalu Cloud Console](https://console.magalu.cloud/).

2. **Create a New Project**:
    - Click on the "Create Project" button.
    - Enter the required project details and click "Submit".

3. **Deploy Your Application**:
    - Navigate to the "Deployments" section.
    - Click on "New Deployment".
    - Upload your application package or connect your repository.
    - Configure the deployment settings as needed.
    - Click "Deploy".

4. **Monitor Your Application**:
    - Go to the "Monitoring" section to view the status and logs of your application.
    - Set up alerts and notifications if necessary.

5. **Manage Resources**:
    - Use the "Resources" section to manage your application's resources such as databases, storage, and networking.

For more detailed instructions, refer to the [Magalu Cloud Documentation](https://docs.magalu.cloud/).
```



install docker via snap 
sudo snap install docker

docker-compose up -d



Em Sistemas Linux ou MacOS
Abra o terminal: A maioria das distribuições Linux já vem com o OpenSSH instalado, então você pode usar o terminal diretamente.

Gerar a chave SSH:

Execute o seguinte comando:
ssh-keygen -t rsa

Você será solicitado a escolher um local para salvar a chave. Pressione Enter para aceitar o local padrão (~/.ssh/id_rsa).
Em seguida, será solicitado a criar uma senha para proteger a chave privada (opcional, mas recomendado).
Verifique a chave SSH: Sua chave pública estará disponível em ~/.ssh/id_rsa.pub. Para visualizar, use:
 cat ~/.ssh/id_rsa.pub