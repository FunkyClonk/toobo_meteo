# Toobo — Bot Météo Telegram

Salut les amis ! C'est moi, **Toobo** 👋

![Toobo et la météo](img/toobo.png)

Je suis un bot Telegram qui t'envoie la météo du jour **tous les jours à 7h**, avec des conseils vestimentaires pour que tu ne sors jamais sans le bon équipement !
*(Mon créateur a la fâcheuse tendance d'oublier son parapluie les jours de pluie...)*

---

## Utiliser la version déployée

Tu peux utiliser ma version déployée directement en m'envoyant un message Telegram avec la commande `/add <prénom> <ville>` (par exemple `/add Toobo Paris`) :

![QR Code Telegram](img/qrcodetelgram.png)

---

## Déployer ton propre Toobot

### 1. Créer son bot Telegram

Toobot est un bot Telegram. Tu devras donc :
1. Créer un compte Telegram si ce n'est pas déjà fait
2. Envoyer un message à **@BotFather**
3. Envoyer la commande `/newbot`
4. Choisir un nom pour ton Toobot
5. BotFather te renverra un **token API** — garde-le précieusement, on en aura besoin très vite !

---

### 2. Choisir sa cible

Toobot peut envoyer ses messages à un groupe ou à un utilisateur directement. Pour cela :

1. Envoie `/add <prénom> <ville>` dans la conversation souhaitée (groupe ou privé)
2. Remplace `TOKEN_API_TOOBOT` par ton token personnel et visite :
```
https://api.telegram.org/botTOKEN_API_TOOBOT/getUpdates
```
3. Tu y trouveras un champ de cette forme : `"chat":{"id":-CLIENT_ID`

Tu as maintenant ton `CLIENT_ID` — garde-le bien au chaud !

---

### 3. Déploiement de la base de données

Toobot tourne actuellement avec une table **Supabase**. La structure attendue est la suivante :

```sql
CREATE TABLE chat_info_telegram_toobo (
    id        BIGSERIAL PRIMARY KEY,
    chat_id   BIGINT UNIQUE NOT NULL,
    city      TEXT NOT NULL,
    chat_name TEXT NOT NULL
);
```

Crée un projet sur [supabase.com](https://supabase.com), crée cette table, puis récupère ta **connection string** dans *Settings > Database > Connection string*.

---

### 4. Déploiement sous GitHub Actions

1. Fork ce projet ou copie-le indépendamment
2. Clique sur **Settings > Secrets and variables > Actions**
3. Renseigne les variables suivantes :

```
TELEGRAM_TOKEN=TOKEN_API_TOOBOT
DB_PULLER=postgresql://postgres:[PASSWORD]@[HOST]:5432/postgres
```

Et voilà — ton propre Toobot enverra ses conseils météo à toi et tes proches chaque matin !

---

### 5. Test en local

```bash
git clone https://github.com/FunkyClonk/toobo_meteo.git
cd toobo_meteo
cp .env.example .env
vi .env
```

Renseigne les variables avec ce que tu as récupéré dans les étapes précédentes :

```env
TELEGRAM_TOKEN=TOKEN_API_TOOBOT
DB_PULLER=postgresql://postgres:[PASSWORD]@[HOST]:5432/postgres
```

Lance ensuite le bot :

```bash
go run .
```

---

## Stack technique

- **Go** — logique principale
- **Telegram Bot API** — envoi des messages
- **Open-Meteo API** — données météo (gratuit, sans clé)
- **Supabase** — base de données PostgreSQL hébergée
- **GitHub Actions** — automatisation quotidienne à 7h