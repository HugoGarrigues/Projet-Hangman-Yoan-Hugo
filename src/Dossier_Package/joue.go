package Dossier_Package

func (j *Joueur)lancement(){
	mot := motAleatoire()
	mot_cache := motCache(mot)
	for !motEstTrouve(mot_cache) && j.essais != 0{
		afficheMotCache(mot_cache)
		lettre := saisieLettre()
		j.chaine_mot = append(j.chaine_mot, lettre)
		if lettreEstPresente(lettre, mot) {
			mot_cache = afficheMotAvecLettreTrouvee(lettre, mot, mot_cache)
		}else {
			j.essais--
		}
	}
	afficheResultat(mot_cache, mot,j.essais)
}