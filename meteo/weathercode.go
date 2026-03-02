package meteo

import "fmt"

func getWeatherCodeTraduction(weatherCode int) string {
	fmt.Println(weatherCode)
	var weatherSignification string
	switch weatherCode {
	case 0:
		weatherSignification = "Développement nuageux non observé ou non observable"
	case 1:
		weatherSignification = "Nuages généralement en dissolution ou devenant moins développés"
	case 2:
		weatherSignification = "État du ciel globalement inchangé"
	case 3:
		weatherSignification = "Nuages généralement en formation ou en développement"
	case 4:
		weatherSignification = "Visibilité réduite par la fumée (incendies, fumée industrielle, cendres volcaniques)"
	case 5:
		weatherSignification = "Brume sèche"
	case 6:
		weatherSignification = "Poussière en suspension généralisée, non soulevée par le vent à la station"
	case 7:
		weatherSignification = "Poussière ou sable soulevé par le vent, sans tempête ni tourbillons développés"
	case 8:
		weatherSignification = "Tourbillons de poussière ou de sable observés, sans tempête"
	case 9:
		weatherSignification = "Tempête de poussière ou de sable en vue"
	case 10:
		weatherSignification = "Brume légère"
	case 11:
		weatherSignification = "Bancs de brouillard ou brouillard givrant peu épais"
	case 12:
		weatherSignification = "Brouillard ou brouillard givrant continu peu épais"
	case 13:
		weatherSignification = "Éclairs visibles, sans tonnerre"
	case 14:
		weatherSignification = "Précipitations visibles n’atteignant pas le sol"
	case 15:
		weatherSignification = "Précipitations visibles atteignant le sol, mais éloignées (>5 km)"
	case 16:
		weatherSignification = "Précipitations visibles atteignant le sol, proches mais pas à la station"
	case 17:
		weatherSignification = "Orage sans précipitations au moment de l’observation"
	case 18:
		weatherSignification = "Rafales observées"
	case 19:
		weatherSignification = "Nuage en entonnoir (trombe) observé"
	case 20:
		weatherSignification = "Bruine (non verglaçante) ou neige en grains, sans averses"
	case 21:
		weatherSignification = "Pluie (non verglaçante), sans averses"
	case 22:
		weatherSignification = "Neige, sans averses"
	case 23:
		weatherSignification = "Pluie et neige ou grésil, sans averses"
	case 24:
		weatherSignification = "Bruine ou pluie verglaçante, sans averses"
	case 25:
		weatherSignification = "Averses de pluie"
	case 26:
		weatherSignification = "Averses de neige ou pluie et neige"
	case 27:
		weatherSignification = "Averses de grêle ou pluie et grêle"
	case 28:
		weatherSignification = "Brouillard ou brouillard givrant"
	case 29:
		weatherSignification = "Orage (avec ou sans précipitations)"
	case 30:
		weatherSignification = "Tempête de poussière ou de sable faible à modérée, en diminution"
	case 31:
		weatherSignification = "Tempête de poussière ou de sable faible à modérée, sans changement notable"
	case 32:
		weatherSignification = "Tempête de poussière ou de sable faible à modérée, en augmentation"
	case 33:
		weatherSignification = "Tempête de poussière ou de sable forte, en diminution"
	case 34:
		weatherSignification = "Tempête de poussière ou de sable forte, sans changement notable"
	case 35:
		weatherSignification = "Tempête de poussière ou de sable forte, en augmentation"
	case 36:
		weatherSignification = "Chasse-neige faible à modérée, basse"
	case 37:
		weatherSignification = "Fort chasse-neige, bas"
	case 38:
		weatherSignification = "Chasse-neige faible à modérée, élevée"
	case 39:
		weatherSignification = "Fort chasse-neige, élevé"
	case 40:
		weatherSignification = "Brouillard à distance, non présent à la station durant l’heure précédente"
	case 41:
		weatherSignification = "Brouillard par bancs"
	case 42:
		weatherSignification = "Brouillard, ciel visible, en diminution"
	case 43:
		weatherSignification = "Brouillard, ciel non visible, en diminution"
	case 44:
		weatherSignification = "Brouillard, ciel visible, sans changement notable"
	case 45:
		weatherSignification = "Brouillard, ciel non visible, sans changement notable"
	case 46:
		weatherSignification = "Brouillard, ciel visible, en formation ou s’épaississant"
	case 47:
		weatherSignification = "Brouillard, ciel non visible, en formation ou s’épaississant"
	case 48:
		weatherSignification = "Brouillard déposant du givre, ciel visible"
	case 49:
		weatherSignification = "Brouillard déposant du givre, ciel non visible"
	case 50:
		weatherSignification = "Bruine intermittente faible, non verglaçante"
	case 51:
		weatherSignification = "Bruine continue faible, non verglaçante"
	case 52:
		weatherSignification = "Bruine intermittente modérée, non verglaçante"
	case 53:
		weatherSignification = "Bruine continue modérée, non verglaçante"
	case 54:
		weatherSignification = "Bruine intermittente forte, non verglaçante"
	case 55:
		weatherSignification = "Bruine continue forte, non verglaçante"
	case 56:
		weatherSignification = "Bruine verglaçante faible"
	case 57:
		weatherSignification = "Bruine verglaçante modérée ou forte"
	case 58:
		weatherSignification = "Bruine et pluie faibles"
	case 59:
		weatherSignification = "Bruine et pluie modérées ou fortes"
	case 60:
		weatherSignification = "Pluie intermittente faible, non verglaçante"
	case 61:
		weatherSignification = "Pluie continue faible, non verglaçante"
	case 62:
		weatherSignification = "Pluie intermittente modérée, non verglaçante"
	case 63:
		weatherSignification = "Pluie continue modérée, non verglaçante"
	case 64:
		weatherSignification = "Pluie intermittente forte, non verglaçante"
	case 65:
		weatherSignification = "Pluie continue forte, non verglaçante"
	case 66:
		weatherSignification = "Pluie verglaçante faible"
	case 67:
		weatherSignification = "Pluie verglaçante modérée ou forte"
	case 68:
		weatherSignification = "Pluie ou bruine et neige faibles"
	case 69:
		weatherSignification = "Pluie ou bruine et neige modérées ou fortes"
	case 70:
		weatherSignification = "Chute intermittente de flocons de neige faible"
	case 71:
		weatherSignification = "Chute continue de flocons de neige faible"
	case 72:
		weatherSignification = "Chute intermittente de flocons de neige modérée"
	case 73:
		weatherSignification = "Chute continue de flocons de neige modérée"
	case 74:
		weatherSignification = "Chute intermittente de flocons de neige forte"
	case 75:
		weatherSignification = "Chute continue de flocons de neige forte"
	case 76:
		weatherSignification = "Poussière de diamant (avec ou sans brouillard)"
	case 77:
		weatherSignification = "Neige en grains"
	case 78:
		weatherSignification = "Cristaux de neige isolés en forme d’étoile"
	case 79:
		weatherSignification = "Granules de glace"
	case 80:
		weatherSignification = "Averses de pluie faibles"
	case 81:
		weatherSignification = "Averses de pluie modérées ou fortes"
	case 82:
		weatherSignification = "Averses de pluie violentes"
	case 83:
		weatherSignification = "Averses de pluie et neige mêlées faibles"
	case 84:
		weatherSignification = "Averses de pluie et neige mêlées modérées ou fortes"
	case 85:
		weatherSignification = "Averses de neige faibles"
	case 86:
		weatherSignification = "Averses de neige modérées ou fortes"
	case 87:
		weatherSignification = "Averses de grésil ou petite grêle faibles"
	case 88:
		weatherSignification = "Averses de grésil ou petite grêle modérées ou fortes"
	case 89:
		weatherSignification = "Averses de grêle faibles, sans orage"
	case 90:
		weatherSignification = "Averses de grêle modérées ou fortes, sans orage"
	case 91:
		weatherSignification = "Pluie faible avec orage durant l’heure précédente"
	case 92:
		weatherSignification = "Pluie modérée ou forte avec orage durant l’heure précédente"
	case 93:
		weatherSignification = "Neige ou pluie et neige mêlées faibles avec orage durant l’heure précédente"
	case 94:
		weatherSignification = "Neige ou pluie et neige mêlées modérées ou fortes avec orage durant l’heure précédente"
	case 95:
		weatherSignification = "Orage faible ou modéré avec pluie et/ou neige"
	case 96:
		weatherSignification = "Orage faible ou modéré avec grêle"
	case 97:
		weatherSignification = "Orage fort avec pluie et/ou neige"
	case 98:
		weatherSignification = "Orage avec tempête de poussière ou de sable"
	case 99:
		weatherSignification = "Orage fort avec grêle"
	default:
		weatherSignification = "Code météo inconnu"
	}
	return weatherSignification
}
