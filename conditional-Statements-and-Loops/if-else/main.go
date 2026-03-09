package ifelse

// In Go’s philosophy, it is better to avoid unnecessary branches and indentation of code. It is also considered better to return as early as possible

// Parameters
// Parameter	Description
// role	"admin", "engineer", "intern", "visitor"
// hour	0–23 (24 hour format)
// clearance	security clearance level (1–5)
// emergency	emergency override
// maintenance	system maintenance ongoing
func CanEnter(role string, hour int, clearance int, emergency bool, maintenance bool) string {

	// Input validation
	if hour < 0 || hour > 23 {
		return "INVALID TIME"
	}

	if role != "admin" && role != "engineer" && role != "intern" && role != "visitor" {
		return "INVALID ROLE"
	}

	if clearance < 1 || clearance > 5 {
		return "INVALID CLEARANCE"
	}

	// Emergency rule
	if emergency {
		if role == "admin" || role == "engineer" {
			return "ACCESS GRANTED: Emergency Override"
		}
		return "ACCESS DENIED"
	}

	// Maintenance rule
	if maintenance {
		if role == "admin" && clearance >= 4 {
			return "ACCESS GRANTED: Maintenance Access"
		}
		return "ACCESS DENIED"
	}

	// After hours rule
	if hour < 8 || hour > 18 {
		if role == "admin" {
			return "ACCESS GRANTED: Admin After Hours"
		}
		return "ACCESS DENIED"
	}

	// Normal hours role rules
	switch role {

	case "admin":
		if clearance >= 3 {
			return "ACCESS GRANTED: Admin Access"
		}

	case "engineer":
		if clearance >= 2 {
			return "ACCESS GRANTED: Engineer Access"
		}

	case "intern":
		if clearance >= 1 && hour >= 10 && hour <= 16 {
			return "ACCESS GRANTED: Intern Access"
		}

	case "visitor":
		if hour >= 12 && hour <= 15 {
			return "ACCESS GRANTED: Visitor Access"
		}
	}

	return "ACCESS DENIED"
}



func main(){

}