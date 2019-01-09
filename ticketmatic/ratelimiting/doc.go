// Each account has a built-in rate limiter and may optionally use queues to
// streamline the handling of peak sales. These rate limiters are also applicable
// to the API: order creation (and ticket assignment) is limited globally for an
// account, so a high demand on the web sales pages will also impact the available
// capacity on the API (and vice-versa).
//
// This rate limiting can manifest itself when creating orders or when adding
// tickets to an order: a response status of 429 Rate Limit Exceeded will be
// returned.
//
// The rate limiter works with 2 time intervals: per minute and per 5-second
// interval. The allowed rate per 5-second interval is equal to the rate per minute
// / 3. So for example: if the rate limit per minute is 60, the rate limit per 5
// seconds is 20.
//
// Rate limiting in libraries
//
// Go
//
// A rate limited call will return a RateLimitError. This error has a Backoff field
// which contains the number of seconds to wait before making a new request.
//
//
//    _, err = orders.Create(c, &ticketmatic.CreateOrder{
//        Events: []int64{
//            777714,
//        },
//        Saleschannelid: 1,
//    })
//    if err != nil {
//        if e, ok := err.(*ticketmatic.RateLimitError); ok {
//            // Do something useful with e:
//            return fmt.Errorf("Need to sleep for %d seconds\n", e.Backoff)
//        } else {
//            return err
//        }
//    }
//
//
//
// PHP
//
// A rate limited call will throw a RateLimitException. This exception has a
// backoff field which contains the number of seconds to wait before making a new
// request.
//
//
//    try {
//        $order = Orders::create($client, array(
//            "events" => array(
//                777714,
//            ),
//            "saleschannelid" => 1,
//        ));
//    } catch (RateLimitException $ex) {
//        $backoff = $ex->backoff;
//        // Sleep for $backoff seconds before retrying.
//    }
//
//
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/ratelimiting).
package ratelimiting
