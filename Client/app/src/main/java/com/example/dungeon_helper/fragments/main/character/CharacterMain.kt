package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.R
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentCharacterMainBinding
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.SharedViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException

class CharacterMain : Fragment() {

    companion object {
        fun newInstance() = CharacterMain()
    }
    private lateinit var viewModel: CharacterMainViewModel


    private var _binding: FragmentCharacterMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val characterMainViewModel = ViewModelProvider(this)[CharacterMainViewModel::class.java]

        _binding = FragmentCharacterMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textCharacter
        characterMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        var newCharacter = sharedViewModel.newCharacter.value
        sharedViewModel.newCharacter.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            newCharacter = it
        })

        var allChar = sharedViewModel.allChar.value
        sharedViewModel.allChar.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            allChar = it
        })

        var charId = 16

        GlobalScope.launch(Dispatchers.Main) {

            var token: String = ""
            sharedViewModel.token.observe(viewLifecycleOwner, Observer {
                // updating data in displayMsg
                token = it
            })


            //println(getCookieValue())

            val client = OkHttpClient.Builder().addInterceptor { chain ->
                val original = chain.request()
                val authorized = original.newBuilder()
                    .header("Authorization", token)
                    .build()
                chain.proceed(authorized)
            }
                .build()

            val request = Request.Builder()
                .url("http://194.247.187.44:5000/character/getAllCharactersByAccId")
                .get()
                .build()

            try {
                val response = withContext(Dispatchers.IO) {
                    client.newCall(request).execute()
                }

                if (!response.isSuccessful) {
                    throw IOException("Запрос к серверу не был успешен:" +
                            " ${response.code} ${response.message}")
                }

                println("${response.code} ${response.message}")

                val deserializedCharacters = JsonHelper.deserializeListFromJsonAllCharacter(response.body!!.string())
                println("Десериализованный объект:\n$deserializedCharacters")

                sharedViewModel.allChar.value = deserializedCharacters

            } catch (e: IOException) {
                println("Ошибка подключения: $e");
            }
        }

        val createCharBtn = binding.createCharBtn
        createCharBtn.setOnClickListener{
            val activity = requireActivity() as MainActivity
            val navView = activity.getNavView()
            val menu = navView.menu
            val menuItem1 = menu.findItem(R.id.navigation_info)
            val menuItem2 = menu.findItem(R.id.navigation_character)
            val menuItem3 = menu.findItem(R.id.navigation_lobby)
            val menuItem4 = menu.findItem(R.id.navigation_account)
            menuItem1.isVisible = false
            menuItem2.isVisible = false
            menuItem3.isVisible = false
            menuItem4.isVisible = false
            (activity as MainActivity).navController.navigate(R.id.action_navigation_character_to_characterCreate1)
        }

        val characterView = binding.characterView
        characterView.setOnClickListener{
            GlobalScope.launch(Dispatchers.Main) {


                val jsonBody = JSONObject().apply {
                    put("id", allChar?.get(1)?.idChar)
                }


                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                println(body)

                var token: String = ""
                sharedViewModel.token.observe(viewLifecycleOwner, Observer {
                    // updating data in displayMsg
                    token = it
                })


                //println(getCookieValue())

                val client = OkHttpClient.Builder().addInterceptor { chain ->
                    val original = chain.request()
                    val authorized = original.newBuilder()
                        .header("Authorization", token)
                        .build()
                    chain.proceed(authorized)
                }
                    .build()

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/character/getCharacterById")
                    .post(body)
                    .build()

                try {
                    val response = withContext(Dispatchers.IO) {
                        client.newCall(request).execute()
                    }

                    if (!response.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response.code} ${response.message}")
                    }

                    println("${response.code} ${response.message}")


                    val deserializedCharacter = JsonHelper.deserializeFromJsonGetCharacter(response.body!!.string())
                    println("Десериализованный объект:\n$deserializedCharacter")

                    sharedViewModel.getCharacter.value = deserializedCharacter

                    val activity = requireActivity() as MainActivity
                    val navView = activity.getNavView()
                    val menu = navView.menu
                    val menuItem1 = menu.findItem(R.id.navigation_info)
                    val menuItem2 = menu.findItem(R.id.navigation_character)
                    val menuItem3 = menu.findItem(R.id.navigation_lobby)
                    val menuItem4 = menu.findItem(R.id.navigation_account)
                    menuItem1.isVisible = false
                    menuItem2.isVisible = false
                    menuItem3.isVisible = false
                    menuItem4.isVisible = false
                    (activity as MainActivity).navController.navigate(R.id.action_navigation_character_to_characterViewing)

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }



        }

    }

}