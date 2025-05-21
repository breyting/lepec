# Lepec

## Présentation du projet

Projet de logiciel éducatif python (ou de programmation de manière générale) en classe

L'objectif de ce projet est de réaliser un logiciel en ligne permettant d'apprendre python (ou un autre langage) en classe

Les principales fonctionnalités sont :

- un interpréteur python intégré dans le navigateur
- une fenêtre permettant de partager un pdf ou un cours créer sur le site
- avoir une session partagée avec les étudiants afin de pouvoir corriger le code ensemble au tableau avec la session professeur

Autres points de développement plus impactant :

- Utiliser les comptes utilisateurs du domaine AD des professeurs pour créer des sessions (permettra d'éviter de créer des sessions infiniment)
- Les étudiants se connecterait avec un code de session (un peu comme wooclap) généré lorsque le professeur créera une session.
  - Les étudiants pourraient se connecter également avec leur compte AD pour sauvegarder leur code ? Permettrait de les connecter un peu plus et de connecter les élèves.
- Créer une interface d'accueil
- Sauvegarde des données utilisateurs ? comment ? Sauvegarde des fichiers ?
  -> permettrait de télécharger le code rédiger, mais pas de sauvegarder dans une DB. rendrait l'application plus légère.
- Lister les interactions possible professeur/élève :
  -> professeur : consulter le code que les élèves soumettent. élèves auraient 2 boutons : "run" et "submit"

## logiciels nécessaires

Ici se trouve la liste des logiciels nécessaire pour le bon fonctionnement du site.

Docker
Python
Golang
