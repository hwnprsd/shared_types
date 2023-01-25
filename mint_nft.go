package shared_types

type MintPoapMessage struct {
	MessagingBase
	Email    string
	Address  string
	Name     string
	TokenURI string
}

func NewMintPoapMessage(email string, walletAddress string, name string, tokenUri string) *MintPoapMessage {
	return &MintPoapMessage{
		MessagingBase: MessagingBase{
			WorkType: "MINT_POAP",
		},
		Name:     name,
		Email:    email,
		Address:  walletAddress,
		TokenURI: tokenUri,
	}
}

type MintQuizNFTMessage struct {
	MessagingBase
	Address  string
	TokenURI string
	Email    string
}

func NewMintQuizNFTMessage(email string, walletAddress string, tokenUri string) *MintQuizNFTMessage {
	return &MintQuizNFTMessage{
		MessagingBase: MessagingBase{
			WorkType: "MINT_QUIZ_NFT",
		},
		Email:    email,
		Address:  walletAddress,
		TokenURI: tokenUri,
	}
}
