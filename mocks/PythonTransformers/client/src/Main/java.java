/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package Main;



/**
 *
 * @author jcaballero
 */
import java.util.concurrent.TimeUnit;
import org.mockserver.client.server.MockServerClient;
import org.mockserver.matchers.Times;
import static org.mockserver.model.HttpRequest.request;
import org.mockserver.model.Delay;
import org.mockserver.model.Header;
import static org.mockserver.model.HttpResponse.response;
public class java {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        new MockServerClient("localhost", 5000)
                .when(
                        request()
                                .withMethod("POST")
                                .withPath("/api/builds"),
                        Times.exactly(1)
                )
                .respond(
                        response()
                                .withStatusCode(201)
                                .withHeaders(
                                        new Header("Content-Type", "application/json; charset=utf-8")
                                )
                                .withBody("{ \"data\": \"v0.2.2-beta_1\" }")
                                .withDelay(new Delay(TimeUnit.SECONDS, 0))
                );

        new MockServerClient("localhost", 5000)
                .when(
                        request()
                                .withMethod("POST")
                                .withPath("/api/deploys"),
                        Times.exactly(1)
                )
                .respond(
                        response()
                                .withStatusCode(201)
                                .withHeaders(
                                        new Header("Content-Type", "application/json; charset=utf-8")
                                )
                                .withBody("{\"message\": \"Deploy success\"}")
                                .withDelay(new Delay(TimeUnit.SECONDS, 0))
                );

        new MockServerClient("localhost", 5000)
                .when(
                        request()
                                .withMethod("GET")
                                .withPath("/api/builds"),
                        Times.exactly(1)
                )
                .respond(
                        response()
                                .withStatusCode(201)
                                .withHeaders(
                                        new Header("Content-Type", "application/json; charset=utf-8")
                                )
                                .withBody("{ \"beta\": [ \"v0.0.1-beta_1\", \"v0.0.1-beta_2\", \"v0.0.1-beta_3\", \"v0.0.2-beta_1\", \"v0.0.3-beta_1\", \"v0.1.1-beta_1\", \"v0.2.1-beta_1\" ], \"prod\": [ \"v0.0.1\", \"v0.0.2\", \"v0.0.3\", \"v0.1.1\", \"v0.2.1\" ], \"staging\": [] }")
                                .withDelay(new Delay(TimeUnit.SECONDS, 1))
                );
    }

}
