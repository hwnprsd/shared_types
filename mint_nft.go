package shared_types

const WORK_TYPE_MINT_POAP = "MINT_POAP"

type MintPoapMessage struct {
	MessagingBase
	Email           string
	Address         string
	Name            string
	TokenURI        string
	EmailTemplateId uint
}

func NewMintPoapMessage(taskId uint, email, walletAddress, name, tokenUri string, emailTemplateId uint) *MintPoapMessage {
	return &MintPoapMessage{
		MessagingBase:   *NewMessagingBase(taskId, WORK_TYPE_MINT_POAP),
		Name:            name,
		Email:           email,
		Address:         walletAddress,
		TokenURI:        tokenUri,
		EmailTemplateId: emailTemplateId,
	}
}

const WORK_TYPE_MINT_QUIZ_NFT = "MINT_QUIZ_NFT"

type MintQuizNFTMessage struct {
	MessagingBase
	Address  string
	TokenURI string
	Email    string
}

func NewMintQuizNFTMessage(taskId uint, email string, walletAddress string, tokenUri string) *MintQuizNFTMessage {
	return &MintQuizNFTMessage{
		MessagingBase: *NewMessagingBase(taskId, WORK_TYPE_MINT_QUIZ_NFT),
		Email:         email,
		Address:       walletAddress,
		TokenURI:      tokenUri,
	}
}

const WORK_TYPE_CREATE_SC_WALLET = "CREATE_SC_WALLET"

type CreateSmartContractWallet struct {
	MessagingBase
	Address string
}

func CreateSmartContractWalletMessage(taskId uint, ownerAddress string) *CreateSmartContractWallet {
	return &CreateSmartContractWallet{
		MessagingBase: *NewMessagingBase(taskId, WORK_TYPE_CREATE_SC_WALLET),
		Address:       ownerAddress,
	}
}

const WORK_TYPE_RELAY_META_TX = "RELAY_META_TX"

type RelayMetaTxMessage struct {
	MessagingBase
	Data      string
	From      string
	Gas       string
	Nonce     uint
	To        string
	Value     string
	Signature string
}

func CreateRelayMetaTxMessage(taskId uint, Data, From, Gas, To, Value, Signature string, Nonce uint) *RelayMetaTxMessage {
	return &RelayMetaTxMessage{
		*NewMessagingBase(taskId, WORK_TYPE_RELAY_META_TX),
		Data, From, Gas, Nonce, To, Value, Signature,
	}
}
