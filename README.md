Salut les amis!

C'est moi, Toobo(t),

![Toobo et la méteo de Gulli](img/toobo.png)
Je suis un bot Telegram qui t'enverra la météo de la journée tout les jours à 7h ainsi que des conseils pour que tu n'oublies pas de t'équiper correctement!

(Mon créateur à la facheuse tendance d'oublier de prendre son parapluie les jours de pluies)

Pour m'uiliser tu peux soit, me deployer sur ton propre projet git pour de l'automatisation ou me tester en local.

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
CHAT_ID=CLIENT_ID
METEO_URL=https://api.open-meteo.com/v1/forecast?latitude=10.1111&longitude=1.1111&daily=weather_code,sunrise,sunset,wind_speed_10m_max,apparent_temperature_max,apparent_temperature_min,daylight_duration,sunshine_duration,showers_sum,snowfall_sum,precipitation_hours,rain_sum,uv_index_max&models=meteofrance_seamless&timezone=Europe%2FBerlin&forecast_days=1 #TODO: personnaliser en fonction de sa ville
```

Et voila, maintenant tu auras aussi ton propre Toobot qui pourra envoyer ses conseils météo à toi et tes proches

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