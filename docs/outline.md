## services overview

- profiles:
	Description:
		the profiles service should handle the transactions of profile creation, updating and deletion. 
	RPCs:
		- get profile from id
		- update profile from id
		- create profile
	Expected Down:
		profiles is down, user cannot be validated, attempt to continute as if the user access level is 0-> anon user. Also, when
		the output is resolved; attach a statement notifing the user that he/she cannot be found in the system. 

		Another option could be checking the memory records for said user, to determine if they are valid in the system. But this
		would require us to include the user access level in the memory files.		
	Fields:
		### Profile Field
		- Type:	Type of profile (user or device)
		- AccessLevel: init from 0-4
			- anon:		new to the system, needs to be registerd
			- normal:	basic user access (for regular profiles)
			- admin:	admin access to prod systems
			- dev: 		development access to non-prod systems
			- root:		access to all systems	  
		- Name:	name on profile.
		- ID:   id of the profile
	Misc:
		[x]	cached
		[] 	support restful

- memory:
	Description:	
		the memory service should handle the transactions of pulling memory files from users, devices. As well as creating new
		memory files. This service is immutable.
	RPCs:
		- get latest from user
		- get latest from device
		- create memory
	Expected Down:
		if this services goes down memory files cannot be pulled, meaning that some of the skills that require memory files, can no
		longer use them. So, in these services, we should just say something along the lines of "sorry, but i am having trouble rea-
		ding from memory right now". 	
	Misc:	
		[]	cached
		[]	support restful

- scheduler
	Description:
		schedules jobs within the orva network, this should be able to schedule restful calls.
	RPCs:
		- create job
		- get last time job was completed
		- optional: eta for job
	Expected Down:
		if this service goes down, it will not affect the user experience in any way other than show stale data within the system. This
		isn't a huge deal, but as an added feature we should be able to check the last time the job was done, if asked by the user. 
		
		Example:
		user-> "what is the new"; orva-> "gives some news"; user-> "how old is this"; orva->checks date on last job.
	Misc:
		[]	cached
		[] 	support restful

- skills
	Description:
		handles CRUD operations of skills as well as determining what skill is being used. I have a few methods to determining, i want to
		re-implement the old way of doing it as well as a ML method. 
	RPCs:	
		- create, update & delete skill
		- determine skill	
	Expected Down:
		so.. if this is down then the user will not be able to route to determine which skill to use, so we should fall back on just
		the speech service. 

		now in the case of the crud operations not being able to resolve, just return that the service is not active, so try again later.
- speech:
	Description:
		handles input and determines speech. currently without the optional rpcs this service should only be called when a skill does not get 		     detected. 
	RPCs:	
		- determine speech
		- optional: createVariant

	Expected Down:
		when the speech service goes down, we should return a message to basically say that we don't know how to reply to the users input. or 		    we should just do nothing due to the content check later in the pipeline.
		
- core:
	Description:	
		handles statement processing, comibining most of the other servies in determining output. 	

		ideal route:
			input entry-> compile user & device report -> concurrently check for skill-> route to skill with the report/ skillid, perform skill, return data found from skill. 
		
		bad user & device:
			
	RPCs:
		- process statement
	Expected Down:
		if this service is down, then the whole service will not work. Basically just send a message to the user that the core service
		cannot be reached and try again.

		In the future, i would like to implement most of this service as a stand-alone system with embeded skill & speech routines.


