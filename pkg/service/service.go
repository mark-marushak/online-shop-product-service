package service

/*

Service drive every logic inside, what business logic for Auth Service.
1. Register
2. Auth (Basic login and password, OAuth) + 2FA
3. JWT token to other services to control policy and view

More about Auth Flows:
@ service control way to Identify User and get permission to communicate with services
Identity Management
- Multi-Factor Authentication
- Password Management
- Single Sign On
- Two-Factor Authentication

@ service store permission to services
Access Management
depend on  services:
- Credential Management - take
- Data Security - take permission

@ service store permission to credentials
Credential Management (CRUD)
relate to Access Management - return permission on doing action
depend on Data Security and User Management - check if user who want create credential has permission t


@ service store users in simple meaning
 another services responsible for permission and access for this user
User Management (CRUD)
relate to Identity Management
relate to Access Management
relate to Data Security
*/

type Service struct {
}

func New() error {
	return nil
}

func (s *Service) Token() error {
	return nil
}

func (s *Service) CreateCredential()
