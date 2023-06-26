package app
import ("testing")

func testIsEmail(t *testing.T){
	_, err := IsEmail("hello")
	if err == nil{
		t.Error("Email is invalid")
	}

	_, err = IsEmail("test@gmail.com")
	if err == nil{
		t.Error("Email is valid")
	}
}