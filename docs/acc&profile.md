## This outline differs the account from the profile.

Accounts only should hold information about the specific account. Doesn't matter if it is a phone, user or iot device they are all "accounts".
After accounts are created, accounts are then given an account type. Accounts also hold "safe" data that does not get exposed to the skill services. 
So things like passwords, account access, etc should be held in the account information and NOT in the profiles.

Profiles are a user only model that hold specific information about a said user. This is basically what gets exposed to the skil services. Examples of what could be held in the 
profiles are "first name", "last name", "favorite color", etc.

