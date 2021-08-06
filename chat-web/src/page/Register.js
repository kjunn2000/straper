import axios from '../axios/api'
import React,{useState} from 'react'
import { Link, useHistory } from 'react-router-dom'

const Register = () => {

	const history = useHistory()

	const [RegisterForm, setRegisterForm] = useState({username:"", password:"",email:"",phone_no:""})

	const updateForm = (e) => {
		setRegisterForm({...RegisterForm,[e.target.name]:e.target.value})
	}

	const onRegister= () => {
		axios.post("http://localhost:8080/api/v1/account/opening",RegisterForm)
			.then(res => {
				if(res.data.Success){
					history.push("/login")
				}
			})
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
				<div className="self-center">
					<div>Email</div>
					<input type="email" className="bg-gray-800 p-2 rounded-lg" name="email" onChange={(e)=>updateForm(e)}/>
				</div>
				<div className="self-center">
					<div>Phone No</div>
					<input type="number" className="bg-gray-800 p-2 rounded-lg" name="phone_no" onChange={(e)=>updateForm(e)}/>
				</div>
				<button type="button" className="bg-indigo-400 self-center w-48 p-1" onClick={()=>onRegister()}>
					REGISTER NOW
				</button>
				<Link to="/login" className="text-indigo-300 self-center cursor-pointer hover:text-indigo-500">Go to login ?</Link>
			</form>
		</div>
	)
}

export default Register 