# Gorch

## Ovierview

Gorch è un orchestratore che trae ispirazione da **systemd** e **docker-compose**.

- **docker-compose** è un ottimo orchestratore ma è troppo vicino a docker.
- **systemd** è abbastanza generico ma il modo di configurare i processi e ostico.

Voglio uno strumento che mi permetta di orchestrare facilmente dei processi e che permetta di descriverli facilmente
in un file di configurazione.

Immaginando di lavorare ad un ambiente di sviluppo. In un ambiente di sviluppo è necessario automatizzare
le configurazioni per l'esecuzione di diversi programmi.
Per questo sembra perfetto **docker-compose**. Il problema è che a volte non si vuole lavorare con una astrazione come **docker**
ma si preferisce lavorare direttamente con gli eseguibili che sono installati sul pc.

**Systemd** è perfetto per lo scopo se non fosse per il sistema di configurazione che è incredibilmente ostico.
Avere file di configurazione in una cartella specifica del proprio computer è un ostacolo alla versatilità della configurazione.

La mia idea è di poter creare un file di configurazione tipo **yaml** nella root di un progetto.

Il file di configurazione va a descrivere valori specifici del processo come:

- Variabili d'ambiente
- Working directory
- arguments
- stdin
- stdout, stderr
- logs

Traendo ispirazione da docker immagino alcune delle seguenti feature:

- **run** - Esegue un programma direttamente collegato al terminale corrente.
- **start** - Esegue un programma in background che rimane collegato al demone.
- **restart** - Restart un programma collegato al demone.
- **stop** - Stop un programma collegato al demone.
- **ls** - Elenco dei programmi che possono essere eseguiti.
- **ps** - Lo stato dei programmi in running.
- **serve** - Avvia l'esecuzione in maniera Daemon.

Mi vengono in mente 3 modalità del programma.

- **Runner** - Il programma può eseguire direttamente un programma
- **Daemon** - Il programma controlla l'esecuzione di altri programmi, si mette in ascolto del client.
- **Client** - Il programma invia i comandi al Daemon per recuperare informazioni o eseguire programmi.
