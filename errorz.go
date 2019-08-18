////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package errorz

import (
	"errors"
)

const ErrorsSeparator = "; "

// Combines two Errors into a single Error.
func Combine(
	error1 error,
	error2 error,
) error {

	if error1 == nil {
		if error2 == nil {
			return nil
		} else {
			return error2
		}
	} else {
		if error2 == nil {
			return error1
		} else {
			return errors.New(error1.Error() + ErrorsSeparator + error2.Error())
		}
	}
}
