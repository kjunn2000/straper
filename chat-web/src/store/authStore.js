import create from "zustand"

const useAuthStore = create(set => ({
	accessToken : "",
	setAccessToken: (accessToken) => {
		set((state) => ({
			accessToken: accessToken
		}))
	}
}))

export default useAuthStore