# Cahier des charges

- [Cahier des charges](#cahier-des-charges)
  - [Workers](#workers)
  - [Schedulers / Queues](#schedulers--queues)
  - [Back / API](#back--api)
  - [Front](#front)

## Workers

- Se connecter au compte AWS + SCW de l' utilisateur
- Récupérer les informations liés au compte utilisateur
  - liste des services actifs
  - ...
- Mettre à jour la BDD

## Schedulers / Queues

- Déclencher le/les workers regulièrement
- ajouter des jobs de manière manuelle (merci)


## Back / API

- Stocker de manière sécurisée les token / moyen de connection récurrente
  - openssl key pair
- Récupérer les informations conćernant les services actifs du client + stocker en BDD
- Authentifier l'utilisateur

## Front

- Vue: connection
- Vue: configuration
- Vue: dashboard
- Vue: gestion de compte
- Possibilité de trigger un refresh