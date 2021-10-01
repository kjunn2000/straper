import React ,{Fragment,useState} from 'react'
import ChannelSidebar from './ChannelSidebar';
import SidebarIcon from './SidebarIcon';
import useWorkspaceStore from '../store/workspaceStore';
import { Dialog, Transition } from '@headlessui/react'
import axios from '../axios/api';

function Sidebar(){

	const workspaces = useWorkspaceStore(state => state.workspaces)
	const addWorkspace = useWorkspaceStore(state => state.addWorkspace)
	const currWorkspace = useWorkspaceStore(state => state.currWorkspace)
	const setCurrWorkspace = useWorkspaceStore(state => state.setCurrWorkspace)
	const [isDialogOpen,setDialogOpen] = useState(false)
	const [addWorkspaceForm, setAddWorkspaceForm] = useState({workspace_name:""})
	const [joinWorkspaceForm, setJoinWorksapceForm] = useState({workspace_id:""})
	const [isAddWorkspaceDialogOpen, setAddWorkspaceDialogOpen] = useState(false)
	const [isJoinWorkspaceDialogOpen, setJoinWorkspaceDialogOpen] = useState(false)

	const changeWorkspace = (workspaceId) => {
		setCurrWorkspace(workspaces.find(workspace => workspace.workspace_id == workspaceId));
	}

	const closeAddWorkspaceModal = () => {
		setDialogOpen(false)
		setAddWorkspaceDialogOpen(false)
	}

	const openAppWorkspaceModal = () => {
		setDialogOpen(true)
		setAddWorkspaceDialogOpen(true)
	}

	const closeJoinWorkspaceModal = () => {
		setDialogOpen(false)
		setJoinWorkspaceDialogOpen(false)
	}

	const updateAddWorkspaceForm =(e)=> {
		setAddWorkspaceForm({...addWorkspaceForm,[e.target.name]:e.target.value});
	}

	const udpateJoinWorkspaceForm =(e)=> {
		setJoinWorksapceForm({...joinWorkspaceForm,[e.target.name]:e.target.value})
	}

	const toggleDialog = () => {
		if (isAddWorkspaceDialogOpen){
			setAddWorkspaceDialogOpen(false)
			setJoinWorkspaceDialogOpen(true)
		}else {
			setAddWorkspaceDialogOpen(true)
			setJoinWorkspaceDialogOpen(false)
		}
	}

	const addNewWorkspace = () => {
		axios.post("http://localhost:8080/api/v1/protected/workspace/create",addWorkspaceForm)
			.then(res => {
				if(res.data.Success){
					const newWorkspace = res.data.Data
					addWorkspace(newWorkspace)
					setCurrWorkspace(newWorkspace)
				}
			})
		closeAddWorkspaceModal()
	}

	const joinWorkspace= () => {
		axios.post("http://localhost:8080/api/v1/protected/workspace/join",joinWorkspaceForm)
			.then(res => {
				if(res.data.Success){
					const newWorkspace = res.data.Data
					addWorkspace(newWorkspace)
					setCurrWorkspace(newWorkspace)
				}
			})
		closeJoinWorkspaceModal()
	}

	return (
		<div className="flex flex-row">
			<div className="flex flex-col w-24 h-screen p-3 bg-black">
				<button className="text-white text-lg font-medium bg-purple-600 
				hover:bg-purple-900 rounded-lg p-1 mb-6 justify-center">SR</button>
				{
					workspaces && workspaces.map(workspace => (
						<SidebarIcon key={workspace.workspace_id} workspace={workspace} changeWorkspace={changeWorkspace}/>
					))
				}
				<button className="rounded-full text-white text-center h-12 w-12 self-center
				bg-red-500 hover:bg-red-800 mt-3" onClick={()=>openAppWorkspaceModal()}>+</button>

			</div>
			<ChannelSidebar workspace={currWorkspace}/>

			<Transition appear show={isDialogOpen} as={Fragment}>
				<Dialog
					as="div"
					className="fixed inset-0 z-10 overflow-y-auto"
					onClose={isAddWorkspaceDialogOpen ? closeAddWorkspaceModal:closeJoinWorkspaceModal}
				>
					<div className="min-h-screen px-4 text-center">
						<Transition.Child
							as={Fragment}
							enter="ease-out duration-300"
							enterFrom="opacity-0"
							enterTo="opacity-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100"
							leaveTo="opacity-0"
						>
							<Dialog.Overlay className="fixed inset-0" />
						</Transition.Child>

						<span className="inline-block h-screen align-middle" aria-hidden="true" > &#8203; </span>
						{ isAddWorkspaceDialogOpen &&
						<Transition.Child
							as={Fragment}
							enter="ease-out duration-300"
							enterFrom="opacity-0 scale-95"
							enterTo="opacity-100 scale-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100 scale-100"
							leaveTo="opacity-0 scale-95"
						>
							<div className="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl space-y-5">
								<Dialog.Title
								as="h3"
								className="text-lg font-medium leading-6 text-gray-900"
								>
									Create Your Own Workspace
								</Dialog.Title>
								<div className="mt-2">
									<div className="self-center space-y-5">
										<div>New Workspace Name</div>
										<input className="bg-gray-200 p-2 w-full" name="workspace_name" onChange={(e)=>updateAddWorkspaceForm(e)}/>
									</div>
								</div>
								
								<div className="text-indigo-500 self-center cursor-pointer hover:text-indigo-300" onClick={()=>toggleDialog()}>Join a workspace?</div>

								<div className="mt-4 flex justify-end">
									<button
										type="button"
										className="inline-flex justify-center px-4 py-2 text-sm font-medium text-blue-900 bg-purple-300 border border-transparent rounded-md hover:bg-purple-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-purple-500"
										onClick={()=>addNewWorkspace()}
										>
											Add
									</button>
								</div>
							</div>
						</Transition.Child>
						}	
						{isJoinWorkspaceDialogOpen &&
						<Transition.Child
							as={Fragment}
							enter="ease-out duration-300"
							enterFrom="opacity-0 scale-95"
							enterTo="opacity-100 scale-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100 scale-100"
							leaveTo="opacity-0 scale-95"
						>
							<div className="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl space-y-5">
								<Dialog.Title
								as="h3"
								className="text-lg font-medium leading-6 text-gray-900"
								>
									Join a workspace
								</Dialog.Title>
								<div className="mt-2">
									<div className="self-center space-y-5">
										<div>Workspace ID (Invite Link)</div>
										<input className="bg-gray-200 p-2 w-full" name="workspace_id" onChange={(e)=>udpateJoinWorkspaceForm(e)}/>
									</div>
								</div>
								
								<div className="text-indigo-500 self-center cursor-pointer hover:text-indigo-300" onClick={()=>toggleDialog()}>Create new workspace?</div>

								<div className="mt-4 flex justify-end">
									<button
										type="button"
										className="inline-flex justify-center px-4 py-2 text-sm font-medium text-blue-900 bg-purple-300 border border-transparent rounded-md hover:bg-purple-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-purple-500"
										onClick={()=>joinWorkspace()}
										>
											Join	
									</button>
								</div>
							</div>
						</Transition.Child>
						}
					</div>
				</Dialog>
			</Transition>
		</div>
	)
}

export default Sidebar

