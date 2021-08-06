import axios from '../axios/api'
import React,{useState} from 'react'
import { Link, useHistory } from 'react-router-dom'
import useAuthStore from '../store/authStore'
import useIdentifyStore from '../store/identityStore'
import useWorkspaceStore from '../store/workspaceStore'

const Login = () => {

	const history = useHistory()

	const [LoginForm, setLoginForm] = useState({username:"", password:""})

	const setAccessToken = useAuthStore(state => state.setAccessToken)

	const setWorkspaces = useWorkspaceStore(state => state.setWorkspaces)
	const setCurrWorkspace = useWorkspaceStore(state => state.setCurrWorkspace)

	const setIdentity = useIdentifyStore(state => state.setIdentity)

	const [errMsg , setErrMsg ] = useState("")
	
	const updateForm = (e) => {
		setLoginForm({...LoginForm,[e.target.name]:e.target.value})
	}

	const onLogin = () => {
		axios.post("http://localhost:8080/api/v1/auth/login",LoginForm,{withCredentials:true})
			.then(res =>{
				if(res.data.Success){
					setAccessToken(res.data.Data?.access_token)
					setIdentity(res.data.Data?.identity)
					fetchWorkspaceData()
				}else if (res.data.ErrorMessage == "invalid.credential"){
					updateErrMsg("Invalid credenital.")
				}else if (res.data.ErrorMessage == "user.not.found"){
					updateErrMsg("User not found.")
				}
			})
	}

	const fetchWorkspaceData = () => {
		axios.get("http://localhost:8080/api/v1/protected/workspace/list")
			.then(res => {
				if (res.data.Success){
					const workspaces = res.data.Data
					if (workspaces && workspaces.length >0 ){
						setWorkspaces(workspaces)
						setCurrWorkspace(workspaces[0])
					}
					history.push("/workspace")
				}
			})
	}

	const updateErrMsg = (msg) => {
		setErrMsg(msg)
		setTimeout(()=> {
			setErrMsg("")
		}, 5000)
	}

	return (
		<div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
			<form className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center py-5">
				<div className="self-center">
					<div className="text-xl font-medium text-center">
						WELCOME BACK
					</div>
					<div className="self-center text-gray-500 text-center">
						Good To See You Again, Friends!
					</div>
				</div>
				<div className="self-center">
					<div>Username</div>
					<input className="bg-gray-800 p-2 rounded-lg" name="username" onChange={(e)=>updateForm(e)}/>
				</div>
				<div className="self-center">
					<div>Password</div>
					<input type="password" className="bg-gray-800 p-2 rounded-lg" name="password" onChange={(e)=>updateForm(e)}/>
				</div>
				{
					errMsg != "" && 
					<div className="text-red-600 self-center">{errMsg}</div>
				}
				<button type="button" className="bg-indigo-400 self-center w-48 p-1" onClick={()=>onLogin()}>
					LET'S GO
				</button>
				<Link to="/register" className="text-indigo-300 self-center cursor-pointer hover:text-indigo-500">Register an account</Link>
			</form>
		</div>
	)
}

export default Login
