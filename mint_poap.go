package sharedtypes

type MintPoapMessage struct {
	MessagingBase
	Email   string
	Address string
	Name    string
}

func NewMintPoapMessage(email string, walletAddress string, name string) *MintPoapMessage {
	return &MintPoapMessage{
		MessagingBase: MessagingBase{
			WorkType: "MINT_POAP",
		},
		Name:    name,
		Email:   email,
		Address: walletAddress,
	}
}
