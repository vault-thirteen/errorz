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
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_Combine(t *testing.T) {

	var aTest *tester.Test
	var err1 error
	var err2 error
	var errCombined error
	var errCombinedExpected error

	aTest = tester.New(t)

	// Test #1. ++
	err1 = errors.New("1")
	err2 = errors.New("2")
	errCombinedExpected = errors.New("1" + ErrorsSeparator + "2")
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #2. +-
	err1 = errors.New("1")
	err2 = nil
	errCombinedExpected = err1
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #3. -+
	err1 = nil
	err2 = errors.New("2")
	errCombinedExpected = err2
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)

	// Test #4. --
	err1 = nil
	err2 = nil
	errCombinedExpected = nil
	errCombined = Combine(err1, err2)
	aTest.MustBeEqual(errCombined, errCombinedExpected)
}
