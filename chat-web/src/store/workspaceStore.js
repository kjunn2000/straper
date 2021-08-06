import create from "zustand"

const useWorkspaceStore = create(set => ({
	workspaces : [],
	currWorkspace : {},
	setWorkspaces: (workspaces) => {
		set((state) => ({
			workspaces : workspaces
		}))
	},
	addWorkspace: (workspace) => {
		set((state)=>({
			workspaces : [...state.workspaces,workspace]
		}))
	},
	deleteWorkspace: (workspaceId) => {
		set((state) => ({
			workspaces : state.workspaces.filter(workspace => workspace.workspace_id != workspaceId),
		}))
	},
	setCurrWorkspace : (workspace) => {
		set((state) => ({
			currWorkspace : workspace
		}))
	},
	resetCurrWorkspace : () => {
		set((state) => ({
			currWorkspace : state.workspaces.length > 0 ? state.workspaces[0] : {}
		}))
	}
}))

export default useWorkspaceStore 