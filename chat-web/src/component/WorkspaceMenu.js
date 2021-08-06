import { Menu, Transition } from '@headlessui/react'
import axios from '../axios/api'
import { Fragment, useEffect ,useState} from 'react'
import useIdentifyStore from '../store/identityStore'
import useWorkspaceStore from '../store/workspaceStore'

export default function WorkspaceMenu({workspace}) {

  const [isCreator, setIsCreator] = useState(false)
  const identity = useIdentifyStore(state => state.identity)
  const deleteWorkspaceAtStore = useWorkspaceStore(state => state.deleteWorkspace)
  const resetCurrWorkspace = useWorkspaceStore(state => state.resetCurrWorkspace)
	
  useEffect(()=> {
    if (identity.user_id === workspace.creator_id){
      setIsCreator(true)
    }else {
      setIsCreator(false)
    }
  },[workspace])

	const deleteWorkspace = () => {
    axios.post(`http://localhost:8080/api/v1/protected/workspace/delete/${workspace.workspace_id}`)
      .then(res => {
        if (res.data.Success){
          deleteWorkspaceAtStore(workspace.workspace_id)
          resetCurrWorkspace()
        }
      })
	}

	const leaveWorkspace = () => {
    axios.post(`http://localhost:8080/api/v1/protected/workspace/leave/${workspace.workspace_id}`)
      .then(res => {
        if (res.data.Success){
          deleteWorkspaceAtStore(workspace.workspace_id)
          resetCurrWorkspace()
        }
      })
	}

	return (
	<div>
		<Menu as="div" className="relative w-full inline-block text-left">
			<div className="w-full">
			<Menu.Button className="inline-flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-black rounded-md bg-opacity-20 hover:bg-opacity-30 focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75">
				{workspace.workspace_name}
			</Menu.Button>
			</div>
			<Transition
			as={Fragment}
			enter="transition ease-out duration-100"
			enterFrom="transform opacity-0 scale-95"
			enterTo="transform opacity-100 scale-100"
			leave="transition ease-in duration-75"
			leaveFrom="transform opacity-100 scale-100"
			leaveTo="transform opacity-0 scale-95"
			>
			<Menu.Items className="absolute left-0 w-56 m-5 origin-top-right bg-white divide-y divide-gray-100 rounded-md shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
			<div className="px-1 py-1">
       <Menu.Item>
          {({ active }) => (
            <button
              className={`${
                active ? 'text-gray-300 font-medium' : 'text-gray-900'
              } group flex rounded-md items-center w-full px-2 py-2 text-sm`}
              onClick={()=> isCreator ? deleteWorkspace() : leaveWorkspace()}
            >
              {active ? (
                <DeleteActiveIcon
                  className="w-5 h-5 mr-2 text-violet-400"
                  aria-hidden="true"
                />
              ) : (
                <DeleteInactiveIcon
                  className="w-5 h-5 mr-2 text-violet-400"
                  aria-hidden="true"
                />
              )}
              {
                isCreator ? "Delete Workspace" : "Leave Workspace"
              }
            </button>
          )}
        </Menu.Item>   
			</div>
			</Menu.Items>
			</Transition>
		</Menu>
	</div>
  )
}

function DeleteInactiveIcon(props) {
  return (
    <svg
      {...props}
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="5"
        y="6"
        width="10"
        height="10"
        fill="#EDE9FE"
        stroke="#A78BFA"
        strokeWidth="2"
      />
      <path d="M3 6H17" stroke="#A78BFA" strokeWidth="2" />
      <path d="M8 6V4H12V6" stroke="#A78BFA" strokeWidth="2" />
    </svg>
  )
}

function DeleteActiveIcon(props) {
  return (
    <svg
      {...props}
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="5"
        y="6"
        width="10"
        height="10"
        fill="#8B5CF6"
        stroke="#C4B5FD"
        strokeWidth="2"
      />
      <path d="M3 6H17" stroke="#C4B5FD" strokeWidth="2" />
      <path d="M8 6V4H12V6" stroke="#C4B5FD" strokeWidth="2" />
    </svg>
  )
}
