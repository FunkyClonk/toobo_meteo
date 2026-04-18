Salut les amis!

C'est moi, Toobo(t),

![Toobo et la méteo de Gulli](img/toobo.png)
Je suis un bot Telegram qui t'enverra la météo de la journée tout les jours à 7h ainsi que des conseils pour que tu n'oublies pas de t'équiper correctement!

(Mon créateur à la facheuse tendance d'oublier de prendre son parapluie les jours de pluies)

Pour m'uiliser tu peux utiliser ma version déployé en envoyant ton surnom et ta ville (par exemple Toobo Paris)

![QR Code Telegram](img/qrcodetelgram.png)


Tu peux aussi me deployer sur ton propre projet git pour de l'automatisation ou me tester en local.

Mais tout d'abord tu devra crée ta propre copie de Toobot :

## Creer son Toobot :

Toobot est un bot Telegram, tu devras donc crée un compte telegram, envoyer un message à @BotFather puis envoyer /newbot, selectionne ensuite le nom de ton Toobot et il te renverra un token API (garde le précieusement, on en aura besoin très vite)

## Choisir sa cible :

Toobot peut envoyer ses messages à un groupe ou un utilisateur directement, pour ça, il te faut envoyer `/hello @ID_DE_TON_TOOBOT` dans la conversation souhaiter.
Ensuite remplace TOKEN_API_TOOBOT par ton token presonnel et va sur : 
https://api.telegram.org/botTOKEN_API_TOOBOT/getUpdates

Tu y trouvera un champ de cette forme `"chat":{"id":-CLIENT_ID,`

Voila tu as maintenant ton client_ID (celui la aussi garde le bien au chaud)

## Deploiement sous github Action

Fork ce projet ou copie le independament.
Ensuite clique sur Settings > Secrets and variables > Actions
Puis renseigne les variables suivantes qu'on est allé chercher ensemble :
```
TELEGRAM_TOKEN=TOKEN_API_TOOBOT
```

Et voila, maintenant tu auras aussi ton propre Toobot qui pourra envoyer ses conseils météo à toi et tes proches

## Deploiement de la base de donnée

Toobot tourne actuellement avec une table dans Supabase, pour deployer la tienne (coming soon..)

## Test en local

En local, tu peux 
```bash
git clone https://github.com/FunkyClonk/toobo_meteo.git
cd toobo.go
chmod +x toboo.go
cp .env.example .env
vi .env
```
Renseigne ensuite les variables avec ce que tu à pu trouver dans les parties précédentes.

Tu peux maintenant tester ton Toobot en envoyant un message  
```bash
go run .
```