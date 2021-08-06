import create from "zustand"

const useIdentifyStore= create(set => ({
	identity: [],
	setIdentity: (identity) => {
		set((state) => ({
			identity: identity
		}))
	}
}))

export default useIdentifyStore 