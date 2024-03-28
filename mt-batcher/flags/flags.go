package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "MT_BATCHER"

func prefixEnvVar(prefix, suffix string) string {
	return prefix + "_" + suffix
}

var (
	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   "BUILD_ENV",
	}
	MtlNetworkNameFlag = cli.StringFlag{
		Name:     "mtl-network-name",
		Usage:    "mantle network name",
		Required: true,
		EnvVar:   "MTL_NETWORK_NAME",
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	L2MtlRpcFlag = cli.StringFlag{
		Name:     "l2-mtl-rpc",
		Usage:    "HTTP provider URL for L2 execution engine",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "L2_MTL_RPC"),
	}
	DtlClientUrlFlag = cli.StringFlag{
		Name:     "dtl-client-url",
		Usage:    "dtl client url for mt batcher",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DTL_CLIENT_URL"),
	}
	DisperserEndpointFlag = cli.StringFlag{
		Name:     "disperser",
		Usage:    "Endpoint at which disperser is available",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DISPERSER_ENDPOINT"),
	}
	RetrieverSocketFlag = cli.StringFlag{
		Name:     "retriever-socket",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "RETRIEVER_SOCKET"),
	}
	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "CHAIN_ID"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "GRAPH_PROVIDER"),
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:     "private",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "PRIVATE_KEY"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the proposer",
		EnvVar: prefixEnvVar(envVarPrefix, "MNEMONIC"),
	}
	SequencerHDPathFlag = cli.StringFlag{
		Name: "sequencer-hd-path",
		Usage: "The HD path used to derive the sequencer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		EnvVar: prefixEnvVar(envVarPrefix, "SEQUENCER_HD_PATH"),
	}
	FeePrivateKeyFlag = cli.StringFlag{
		Name:     "fee-private",
		Usage:    "Ethereum private key for mt fee",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "FEE_PRIVATE_KEY"),
	}
	FeeMnemonicFlag = cli.StringFlag{
		Name:   "fee-mnemonic",
		Usage:  "The mnemonic used to derive the wallets for the mt fee",
		EnvVar: prefixEnvVar(envVarPrefix, "MNEMONIC"),
	}
	FeeHDPathFlag = cli.StringFlag{
		Name:   "fee-hd-path",
		Usage:  "The HD path used to derive the wallets for the mt fee",
		EnvVar: prefixEnvVar(envVarPrefix, "SEQUENCER_HD_PATH"),
	}
	PassphraseFlag = cli.StringFlag{
		Name:   "passphrase",
		Usage:  "passphrase for the seed generation process to increase the seed's security",
		EnvVar: prefixEnvVar(envVarPrefix, "PASSPHRASE"),
	}
	EigenContractAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_ADDRESS"),
	}
	EigenFeeContractAddressFlag = cli.StringFlag{
		Name:     "eigen-fee-address",
		Usage:    "Address of the datalayr fee contract",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_FEE_ADDRESS"),
	}
	BlockOffsetFlag = cli.Uint64Flag{
		Name:   "block-offset",
		Usage:  "The offset between the eigenda contract start and the L2 geth blocks",
		Value:  1,
		EnvVar: prefixEnvVar(envVarPrefix, "BLOCK_OFFSET"),
	}
	RollUpMinTxnFlag = cli.Uint64Flag{
		Name:   "rollup-min-txn",
		Usage:  "Rollup transaction min transactions for eigen da",
		Value:  500,
		EnvVar: prefixEnvVar(envVarPrefix, "ROLLUP_MIN_TXN"),
	}
	FeeSizeSecFlag = cli.StringFlag{
		Name:   "fee-size-sec",
		Usage:  "Rollup transaction fee size",
		Value:  "102400", //
		EnvVar: prefixEnvVar(envVarPrefix, "FEE_SIZE_SEC"),
	}
	FeePerBytePerTimeFlag = cli.Uint64Flag{
		Name:   "fee-per-byte-time",
		Usage:  "Fee per byte time",
		Value:  1,
		EnvVar: prefixEnvVar(envVarPrefix, "FEE_PER_BYTE_TIME"),
	}
	RollUpMaxSizeFlag = cli.Uint64Flag{
		Name:   "rollup-max-size",
		Usage:  "Rollup transaction max size data for eigen da",
		Value:  31600, // ktz for order is 3000
		EnvVar: prefixEnvVar(envVarPrefix, "ROLLUP_MAX_SIZE"),
	}
	EigenLayerNodeFlag = cli.IntFlag{
		Name:   "eigen-layer-node",
		Usage:  "The offset between the CTC contract start and the L2 geth blocks",
		Value:  4,
		EnvVar: prefixEnvVar(envVarPrefix, "EIGEN_LAYER_NODE"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "RESUBMISSION_TIMEOUT"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "NUM_CONFIRMATIONS"),
	}
	SafeAbortNonceTooLowCountFlag = cli.Uint64Flag{
		Name: "safe-abort-nonce-too-low-count",
		Usage: "Number of ErrNonceTooLow observations required to " +
			"give up on a tx at a particular nonce without receiving " +
			"confirmation",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
	}
	MainWorkerPollIntervalFlag = cli.DurationFlag{
		Name:     "main-worker-poll-interval",
		Usage:    "main worker poll interval",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "MAIN_WORKER_POLL_INTERVAL"),
	}
	CheckerWorkerPollIntervalFlag = cli.DurationFlag{
		Name:     "checker-worker-poll-interval",
		Usage:    "checker worker poll interval",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "CHECKER_WORKER_POLL_INTERVAL"),
	}
	FeeWorkerPollIntervalFlag = cli.DurationFlag{
		Name:     "fee-worker-poll-interval",
		Usage:    "fee worker poll interval",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "FEE_WORKER_POLL_INTERVAL"),
	}
	DataStoreDurationFlag = cli.IntFlag{
		Name:     "duration",
		Usage:    "Duration to store blob",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DATA_STORE_DURATION"),
	}
	DataStoreTimeoutFlag = cli.IntFlag{
		Name:     "timeout",
		Usage:    "Blob timeout",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DATA_STORE_TIMEOUT"),
	}
	EigenDaHttpPortFlag = cli.IntFlag{
		Name:     "eigen-da-http-port",
		Usage:    "Eigen da service port",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "EIGEN_DA_HTTP_PORT"),
	}
	RetrieverTimeoutFlag = cli.DurationFlag{
		Name:   "retriever-timeout",
		Usage:  "retriever timeout",
		Value:  50 * time.Millisecond,
		EnvVar: prefixEnvVar(envVarPrefix, "RETRIEVER_TIMEOUT"),
	}
	MtlBatcherEnableFlag = cli.BoolFlag{
		Name:   "mtl-batch-enable",
		Usage:  "roll data to eigen da enable",
		EnvVar: prefixEnvVar(envVarPrefix, "MTL_BATCHER_ENABLE"),
	}
	DbPathFlag = cli.StringFlag{
		Name:     "db-path",
		Usage:    "Path for level db",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "DB_PATH"),
	}
	CheckerBatchIndexFlag = cli.Uint64Flag{
		Name:   "checker-batch-index",
		Usage:  "The port of the metrics server",
		Value:  7300,
		EnvVar: prefixEnvVar(envVarPrefix, "CHECKER_BATCH_INDEX"),
	}
	CheckerEnableFlag = cli.BoolFlag{
		Name:   "checker-enable",
		Usage:  "checker enable",
		EnvVar: prefixEnvVar(envVarPrefix, "CHECKER_ENABLE"),
	}
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_TERMINAL"),
	}
	FeeModelEnableFlags = cli.BoolFlag{
		Name:   "fee-model-enable",
		Usage:  "fee model enable",
		EnvVar: prefixEnvVar(envVarPrefix, "FEE_MODEL_ENABLE"),
	}
	SentryEnableFlag = cli.BoolFlag{
		Name:   "sentry-enable",
		Usage:  "Whether or not to enable Sentry. If true, sentry-dsn must also be set",
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_ENABLE"),
	}
	SentryDsnFlag = cli.StringFlag{
		Name:   "sentry-dsn",
		Usage:  "Sentry data source name",
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_DSN"),
	}
	PollingDurationFlag = cli.DurationFlag{
		Name:   "polling-duration",
		Usage:  "polling duration for fetch data from da graph node",
		Value:  1200 * time.Millisecond,
		EnvVar: prefixEnvVar(envVarPrefix, "POLLING_DURATION"),
	}
	SentryTraceRateFlag = cli.DurationFlag{
		Name:   "sentry-trace-rate",
		Usage:  "Sentry trace rate",
		Value:  50 * time.Millisecond,
		EnvVar: prefixEnvVar(envVarPrefix, "SENTRY_TRACE_RATE"),
	}
	MetricsServerEnableFlag = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Whether or not to run the embedded metrics server",
		EnvVar: prefixEnvVar(envVarPrefix, "METRICS_SERVER_ENABLE"),
	}
	MetricsHostnameFlag = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "The hostname of the metrics server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar(envVarPrefix, "METRICS_HOSTNAME"),
	}
	MetricsPortFlag = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "The port of the metrics server",
		Value:  7300,
		EnvVar: prefixEnvVar(envVarPrefix, "METRICS_PORT"),
	}
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar(envVarPrefix, "HTTP2_DISABLE"),
	}
	EchoDebugFlag = cli.BoolFlag{
		Name:   "echo-debug",
		Usage:  "Echo log debug",
		EnvVar: prefixEnvVar(envVarPrefix, "ECHO_DEBUG"),
	}
	EnableHsmFlag = cli.BoolFlag{
		Name:   "enable-hsm",
		Usage:  "Enalbe the hsm",
		EnvVar: prefixEnvVar(envVarPrefix, "ENABLE_HSM"),
	}
	HsmAPINameFlag = cli.StringFlag{
		Name:   "hsm-api-name",
		Usage:  "the api name of hsm for mt-batcher",
		EnvVar: prefixEnvVar(envVarPrefix, "HSM_API_NAME"),
	}
	HsmFeeAPINameFlag = cli.StringFlag{
		Name:   "hsm-fee-api-name",
		Usage:  "the api name of hsm for mt-batcher fee address",
		EnvVar: prefixEnvVar(envVarPrefix, "HSM_FEE_API_NAME"),
	}
	HsmAddressFlag = cli.StringFlag{
		Name:   "hsm-address",
		Usage:  "the address of hsm key for mt-batcher",
		EnvVar: prefixEnvVar(envVarPrefix, "HSM_ADDRESS"),
	}

	HsmFeeAddressFlag = cli.StringFlag{
		Name:   "hsm-fee-address",
		Usage:  "the address of hsm key for mt-batcher fee",
		EnvVar: prefixEnvVar(envVarPrefix, "HSM_FEE_ADDRESS"),
	}

	HsmCredenFlag = cli.StringFlag{
		Name:   "hsm-creden",
		Usage:  "the creden of hsm key for mt-batcher",
		EnvVar: prefixEnvVar(envVarPrefix, "HSM_CREDEN"),
	}

	MinTimeoutRollupTxnFlag = cli.Uint64Flag{
		Name:     "min-timeout-rollup-txn",
		Usage:    "minimum of timeout Rollup transaction amount for mantle da",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "MIN_TIMEOUT_ROLLUP_TXN"),
	}

	RollupTimeoutFlag = cli.DurationFlag{
		Name:     "rollup-timeout",
		Usage:    "Duration we will wait before rollup batch transaction to Mantle Da",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "ROLLUP_TIMEOUT"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	MtlNetworkNameFlag,
	L1EthRpcFlag,
	L2MtlRpcFlag,
	DtlClientUrlFlag,
	DisperserEndpointFlag,
	RetrieverSocketFlag,
	EigenDaHttpPortFlag,
	ChainIdFlag,
	GraphProviderFlag,
	PrivateKeyFlag,
	MnemonicFlag,
	FeePrivateKeyFlag,
	FeeMnemonicFlag,
	FeeHDPathFlag,
	PassphraseFlag,
	SequencerHDPathFlag,
	EigenContractAddressFlag,
	EigenFeeContractAddressFlag,
	BlockOffsetFlag,
	MinTimeoutRollupTxnFlag,
	RollupTimeoutFlag,
	RollUpMinTxnFlag,
	RollUpMaxSizeFlag,
	FeeSizeSecFlag,
	FeePerBytePerTimeFlag,
	MainWorkerPollIntervalFlag,
	CheckerWorkerPollIntervalFlag,
	FeeWorkerPollIntervalFlag,
	DataStoreDurationFlag,
	DataStoreTimeoutFlag,
	EigenLayerNodeFlag,
	DbPathFlag,
	CheckerBatchIndexFlag,
	CheckerEnableFlag,
	ResubmissionTimeoutFlag,
	NumConfirmationsFlag,
	FeeModelEnableFlags,
	SafeAbortNonceTooLowCountFlag,
	MtlBatcherEnableFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	SentryEnableFlag,
	SentryDsnFlag,
	SentryTraceRateFlag,
	HTTP2DisableFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	EchoDebugFlag,
	EnableHsmFlag,
	HsmAddressFlag,
	HsmAPINameFlag,
	HsmCredenFlag,
	HsmFeeAddressFlag,
	HsmFeeAPINameFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
