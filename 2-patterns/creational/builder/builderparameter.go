package builder

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		fmt.Println("You don't seem to have received this e-mail from a valid address")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		fmt.Println("You don't seem to be sending this e-mail to a valid address")
	}
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) WithSubject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) WithBody(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

func sendMailImpl(email *email) {
	fmt.Printf("E-mail has been sent: %+v\n", email)
}

// In order to protect our email type so people don't interact directly with it
// We can work with Build Parameters
// This way, our clients will only ever interact with the EmailBuilder, not with the email itself.

// 1. We start with a build type, which is a function that receives a *EmailBuilder.
type build func(*EmailBuilder)

// 2. Whenever we want to send as e-mail, we pass a build type
// Which is basically passing a function that receives a *EmailBuilder
func SendEmail(action build) {

	// 3. Inside the SendEmail function, our EmailBuilder is instantiated
	builder := EmailBuilder{}

	// 4. The received function is executed while pointing to our EmailBuilder
	action(&builder)

	// 5. And finally, our business logic is triggered to effectively send the built email
	sendMailImpl(&builder.email)
}

// So how does this all work in practice?
func BuilderParameter() {

	// 6. We call the SendEmail function and pass our builder method calls as a function inside
	SendEmail(
		// 7. At the moment we call the SendEmail function, our EmailBuilder doesn't exist yet
		// 8. As explained in point 3. above, before the passed function is executed, the EmailBuilder gets instantiated
		func(eb *EmailBuilder) {
			// 9. At this point, the EmailBuilder already exists, so we execute the builder method calls
			eb.
				From("foo@bar.com").
				To("bar@baz.com").
				WithSubject("Meeting").
				WithBody("Hello, do you want to meet?")
		},
	// 10. Now that the function action/build has finished running, the sendMailImpl(&builder.email) is called
	)
}
